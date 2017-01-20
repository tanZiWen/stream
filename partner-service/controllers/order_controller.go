package controllers

import (
	"code.isstream.com/stream/lib"
	"github.com/gin-gonic/gin"
	log "github.com/Sirupsen/logrus"
	"strconv"
	"strings"
	"code.isstream.com/stream/partner-service/service/orderservice"
	"code.isstream.com/stream/partner-service/model"
	"time"
	"code.isstream.com/stream/idg"
	"code.isstream.com/stream/partner-service/service/customerservice"
	"code.isstream.com/stream/domain"
)

type orderForm struct {
	CustomerId int64            `json:"customerId,string"`
	Items      []*orderItem      `json:"items"`
	Comment    string            `json:"comment"`
}

type orderItem struct {
	Id         int64    `json:"id,string"`
	SkuId      int64    `json:"skuId,string"`
	Number     int32    `json:"number"`
	Attributes []*attribute `json:"attributes"`
}

type attribute struct {
	Id    int64    `json:"id,string"`
	Value string    `json:"value"`
}

type OrderController struct {
	lib.BaseController
}

/**
	create order on behalf of customer, oriented to partner users
 */
func (ctr *OrderController) CreateOrderForCustomer(c *gin.Context) {
	log.Debug("create order for customer")
	var err error

	uid, err := ctr.GetUid(c); if err != nil {
		ctr.Unauthorized(c)
		return
	}

	log.Debug("customer id ", c.Param("id"))
	customerId, err := lib.Str2int64(c.Param("id")); if err != nil {
		ctr.BadRequest(c)
		return
	}

	customer, err := customerservice.GetCustomerById(customerId); if err != nil {
		panic(err)
	}

	//TODO authorize admin to create order maybe
	if customer.OwnerId != uid {
		log.Warn("try to create order on behalf of customer illegally")
		ctr.Unauthorized(c)
		return
	}

	var form orderForm
	err = c.Bind(&form); if err != nil {
		log.Error("bad request ", err)
		ctr.BadRequest(c)
		return
	}

	var order model.Order = model.Order{}
	order.Id, err = idg.Id(); if err != nil {
		log.Warn("fail to generate id", err)
		panic(err)
	}

	order.CreatorId = uid
	order.CustomerId = customerId
	now := domain.UtcTime(time.Now())
	order.Crt = now
	order.Lut = now
	order.Items = []*model.OrderItem{}

	for _, item := range form.Items {
		orderItem := model.OrderItem{}
		orderItem.Id, err = idg.Id(); if err != nil {
			panic(err)
		}

		orderItem.OrderId = order.Id
		orderItem.ItemId = item.Id
		orderItem.SkuId = item.SkuId
		orderItem.Number = item.Number
		orderItem.Crt = now
		orderItem.Lut = now

		order.Items = append(order.Items, &orderItem)
	}

	orderservice.CaculateOrderPrice(order)
	err = orderservice.CreateOrder(&order); if err != nil {
		log.Warn("fail to create order", err)
		panic(err)
		return
	}

	ctr.SuccessPair(c, "order", order)
}

/**
	submit my own order, oriented to customers
 */
func (ctr *OrderController) SubmitMyOrder(c *gin.Context) {
	uid, err := ctr.GetUid(c); if err != nil {
		ctr.Unauthorized(c)
		return
	}
	log.Debug("user id:", uid)
}

//get order by user_id
func (ctr *OrderController) GetMyCustomerOrders(c *gin.Context) {
	uid, err := ctr.GetUid(c); if err != nil {
		ctr.Unauthorized(c)
		return
	}

	var page int
	pageStr := c.Query("page")

	log.Debug("user id and pageNo:", uid, pageStr)

	if strings.TrimSpace(pageStr) == "" {
		page = 1
	} else {
		page, err = strconv.Atoi(pageStr); if err != nil {
			log.Error("failed to get page", err)
			panic(err)
		}
	}

	orders, err := orderservice.GetMyCustomerOrders(page, uid); if err != nil {
		log.Error("fail to get orders:", err)
		panic(err)
	}

	for _, order := range orders {
		order.CustomerName = lib.MaskName(order.CustomerName)
	}

	ctr.SuccessPair(c, "orders", orders)
}

/**
	get order detail by order_id
	获取订单详情
 */
func (ctr *OrderController) GetOrderDetail(c *gin.Context) {
	_, err := ctr.GetUid(c); if err != nil {
		ctr.Unauthorized(c)
		return
	}

	log.Debug("order id:", c.Param("id"))
	orderId, err := lib.Str2int64(c.Param("id")); if err != nil {
		ctr.BadRequest(c)
		return
	}

	order, err := orderservice.GetOrderDetail(orderId); if err != nil {
		log.Error("fail to get order detail:", err)
		panic(err)
	}

	order.CustomerName = lib.MaskName(order.CustomerName)

	ctr.SuccessPair(c, "order", order)
}

/**
	get order by customer_id
	获取用户的所有订单
 */
func (ctr *OrderController) GetOrdersByCustomerId(c *gin.Context) {

}