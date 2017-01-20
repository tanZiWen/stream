package orderservice

import (
	"code.isstream.com/stream/partner-service/model"
	"code.isstream.com/stream/db"
	log "github.com/Sirupsen/logrus"
	"errors"
)

const (
	QUERY_PARENT_ORDERS_BY_OWNER = "select o.id as id, o.customer_id as customer_id, c.name as customer_name, " +
		"o.creator_id as creator_id, o.total_number as total_number, o.total_price as total_price, o.paid_amount as paid_amount, " +
		"o.ship_status as ship_status, o.done as done, o.has_sub_order as has_sub_order, o.is_sub_order as is_sub_order, " +
		"o.parent_order_id as parent_order_id, o.partner_id as partner_id, o.crt as crt from orders o, customer c where " +
		"o.customer_id = c.id and c.owner_id = ? and c.status = 1 and o.status = 1 and o.is_sub_order = false " +
		"order by o.crt desc limit 20 offset ?;"

	QUERY_SUB_ORDERS = "select id, customer_id, total_number, total_price, paid_amount, ship_status from orders " +
		"where status = 1 and is_sub_order = true and "

	QUERY_ORDER_ITEMS = "select * from order_item where status = 1 and "

	QUERY_PARENT_ORDER_BY_ORDERID = "select o.id as id, o.customer_id as customer_id, c.name as customer_name, " +
		"c.mobile as customer_mobile, o.creator_id as creator_id, o.total_number as total_number, o.total_price as total_price, " +
		"o.paid_amount as paid_amount, o.ship_status as ship_status, o.done as done, o.has_sub_order as has_sub_order, " +
		"o.is_sub_order as is_sub_order, o.parent_order_id as parent_order_id, o.partner_id as partner_id, o.crt as crt from " +
		"orders o, customer c where o.id = ? and o.customer_id = c.id and c.status = 1 and o.status = 1 and o.is_sub_order = false"

	QUERY_SUB_ORDER = "select id, customer_id, total_number, total_price, paid_amount, ship_status from orders " +
		"where status = 1 and is_sub_order = true and parent_order_id = ?"

)

func GetMyCustomerOrders(page int, id int64) ([]*model.OrderRS, error) {
	var (
		orderRs []*model.OrderRS
		err error
	)
	offset := (page - 1) * 20
	err = db.Engine.Sql(QUERY_PARENT_ORDERS_BY_OWNER, id, offset).Find(&orderRs); if err != nil {
		log.Error("get orders info error:", err)
		return nil, err
	}

	parentOrderIds := []int64{}
	allOrderIds := []int64{}
	orderMap := map[int64]*model.OrderRS{}
	for _, order := range orderRs {
		if order.HasSubOrder {
			parentOrderIds = append(parentOrderIds, order.Id)
		} else {
			allOrderIds = append(allOrderIds, order.Id)
		}
		orderMap[order.Id] = order
	}

	if len(parentOrderIds) > 0 {
		var suborders []*model.OrderRS// = []*model.OrdersRS{}
		sql := QUERY_SUB_ORDERS
		sql = sql + db.InOrEqual("parent_order_id", parentOrderIds) + " ORDER BY crt DESC limit 20 offset ?;"


		err = db.Engine.Sql(sql, offset).Find(&suborders); if err != nil {
			log.Error("fail to query sub orders:", err)
			return nil, err
		}

		for _, suborder := range suborders {
			parentOrder := orderMap[suborder.ParentOrderId]
			if parentOrder == nil || parentOrder.Id == 0 {
				return nil, errors.New("fail to query sub orders")
			}

			parentOrder.SubOrders = append(parentOrder.SubOrders, suborder)
			allOrderIds = append(allOrderIds, suborder.Id)
			orderMap[suborder.Id] = suborder
		}
	}

	if len(allOrderIds) > 0 {
		var items []*model.OrderItem
		sql := QUERY_ORDER_ITEMS
		sql = sql + db.InOrEqual("order_id", allOrderIds) +  " ORDER BY crt DESC limit 20 offset ?;"

		err = db.Engine.Sql(sql, offset).Find(&items); if err != nil {
			log.Error("fail to query order items:", err)
			return nil, err
		}

		for _, item := range items {
			order := orderMap[item.OrderId]
			if order == nil || order.Id == 0 {
				return nil, errors.New("fail to query order items")
			}

			order.Items = append(order.Items, item)
		}
	}

	return orderRs, nil
}

func GetOrderDetail(id int64) (*model.OrderRS, error) {
	orderRs := &model.OrderRS{}

	_, err := db.Engine.Sql(QUERY_PARENT_ORDER_BY_ORDERID, id).Get(orderRs); if err != nil {
		log.Error("query parent order for order detail info error:", err)
		return nil, err
	}

	var parentOrderId int64
	var orderIds []int64
	orderMap := map[int64]*model.OrderRS{}
	if orderRs.HasSubOrder {
		parentOrderId = id
	} else {
		orderIds = append(orderIds, id)
	}
	orderMap[orderRs.Id] = orderRs

	if parentOrderId > 0 {
		var suborders []*model.OrderRS// = []*model.OrdersRS{}

		err = db.Engine.Sql(QUERY_SUB_ORDER, parentOrderId).Find(&suborders); if err != nil {
			log.Error("fail to query sub order for order detail", err)
			return nil, err
		}

		for _, suborder := range suborders {
			parentOrder := orderMap[suborder.ParentOrderId]
			if parentOrder == nil || parentOrder.Id == 0 {
				return nil, errors.New("fail to query sub order for order detail")
			}

			parentOrder.SubOrders = append(parentOrder.SubOrders, suborder)
			orderIds = append(orderIds, suborder.Id)
			orderMap[suborder.Id] = suborder
		}
	}

	if len(orderIds) > 0 {
		var items []*model.OrderItem
		sql := QUERY_ORDER_ITEMS
		sql = sql + db.InOrEqual("order_id", orderIds) +  " ORDER BY crt DESC;"

		err = db.Engine.Sql(sql).Find(&items); if err != nil {
			log.Error("fail to query order items for order detail:", err)
			return nil, err
		}

		for _, item := range items {
			order := orderMap[item.OrderId]
			if order == nil || order.Id == 0 {
				return nil, errors.New("fail to query order items for order detail")
			}

			order.Items = append(order.Items, item)
		}
	}

	return orderRs, nil
}

func CaculateOrderPrice(order model.Order) {
	//TODO
}

func CreateOrder(order *model.Order) error {
	if order == nil {
		return errors.New("try to save nil order")
	}

	session := db.Engine.NewSession()
	err := session.Begin(); if err != nil {
		return err
	}

	_, err = session.Omit("status").Insert(order); if err != nil {
		_ = session.Rollback()
		return err
	}

	//TODO query items and set price, apply promotions maybe later
	for _, item := range order.Items {
		log.Debug("save order item ", *item)
		_, err = session.Omit("status").Insert(item); if err != nil {
			_ = session.Rollback()
			return err
		}
	}

	err = session.Commit(); if err != nil {
		_ = session.Rollback()
		return err
	}

	return err
}
