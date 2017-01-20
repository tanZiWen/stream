package model

import (
	"code.isstream.com/stream/domain"
)

type Customer struct {
	Id          int64            `xorm:"id pk" json:"id,string"`
	Name        string           `xorm:"name" json:"name"`
	FirstName   string           `xorm:"first_name" json:"firstname"`
	LastName    string           `xorm:"last_name" json:"lastname"`
	Gender      int              `xorm:"gender" json:"gender,string"`
	Mobile      string           `xorm:"mobile" json:"mobile"`
	Email       string           `xorm:"email" json:"email"`
	Birthday    string           `xorm:"birthday date" json:"birthday"`
	IdType      int16            `xorm:"id_type" json:"idType"`
	IdNumber    string           `xorm:"id_number" json:"idNumber"`
	WxOpenId    string           `xorm:"wx_openid" json:"wxOpenid"`
	MemberGrade int16            `xorm:"member_grade" json:"memberGrade"`
	OwnerId     int64            `xorm:"owner_id" json:"ownerId,string"`
	LastFeedId  int64            `xorm:"last_feed_id" json:"lastFeedId,string"`
	Crt         domain.UtcTime   `xorm:"crt timestampz created" json:"crt"`
	Lut         domain.UtcTime   `xorm:"lut timestampz updated" json:"-"`
	Status      int16            `xorm:"status" json:"-"`

	Feeds       []*FeedRS        `xorm:"-" json:"feeds"`
}

func (c *Customer) TableName() string {
	return "customer"
}

type Order struct {
	Id            int64        `xorm:"id pk" json:"id,string"`
	CustomerId    int64        `xorm:"customer_id" json:"customerId,string"`
	CreatorId     int64        `xorm:"creator_id" json:"creatorId,string"`
	TotalNumber   int32        `xorm:"total_number" json:"totalNumber"`
	TotalPrice    float32      `xorm:"total_price" json:"totalPrice"`
	PaidAmount    float32      `xorm:"paid_amount" json:"paidAmount"`
	ShipStatus    int16        `xorm:"ship_status" json:"shipStatus"`
	Done          bool         `xorm:"done" json:"done"`
	HasSubOrder   bool         `xorm:"has_sub_order" json:"hasSubOrder"`
	IsSubOrder    bool        `xorm:"is_sub_order" json:"isSubOrder"`
	ParentOrderId int64        `xorm:"parent_order_id" json:"parentOrderId,string"`
	PartnerId     int64        `xorm:"partner_id" json:"partnerId,string"`
	Crt           domain.UtcTime    `xorm:"crt" json:"crt"`
	Lut           domain.UtcTime    `xorm:"lut" json:"-"`
	Status        int16        `xorm:"status" json:"-"`

	SubOrders     []*Order        `xorm:"-"`
	Items         []*OrderItem    `xorm:"-"`
}

type OrderRS struct {
	Id            	int64        	`xorm:"id pk" json:"id,string"`
	CustomerId    	int64        	`xorm:"customer_id" json:"customerId,string"`
	CustomerName  	string        	`xorm:"customer_name" json:"customerName"`
	CustomerMobile 	string	   	`xorm:"customer_mobile" json:"customerMobile"`
	CreatorId     	int64        	`xorm:"creator_id" json:"creatorId,string"`
	TotalNumber   	int32        	`xorm:"total_number" json:"totalNumber"`
	TotalPrice    	float32      	`xorm:"total_price" json:"totalPrice"`
	PaidAmount    	float32      	`xorm:"paid_amount" json:"paidAmount"`
	ShipStatus    	int16        	`xorm:"ship_status" json:"shipStatus"`
	Done          	bool         	`xorm:"done" json:"done"`
	HasSubOrder   	bool         	`xorm:"has_sub_order" json:"hasSubOrder"`
	IsSubOrder    	bool        	`xorm:"is_sub_order" json:"isSubOrder"`
	ParentOrderId 	int64        	`xorm:"parent_order_id" json:"parentOrderId,string"`
	PartnerId     	int64        	`xorm:"partner_id" json:"partnerId,string"`
	Crt           	domain.UtcTime  `xorm:"crt" json:"crt"`
	Lut           	domain.UtcTime  `xorm:"lut" json:"-"`
	Status        	int16        	`xorm:"status" json:"-"`

	SubOrders     []*OrderRS        `xorm:"-"`
	Items         []*OrderItem    	`xorm:"-"`
}

func (c *Order) TableName() string {
	return "orders"
}

type OrderItem struct {
	Id            int64        `xorm:"id pk" json:"id,string"`
	ParentOrderId int64        `xorm:"parent_order_id" json:"parentOrderId,string"`
	OrderId       int64        `xorm:"order_id" json:"orderId,string"`
	ItemId        int64        `xorm:"item_id" json:"itemId,string"`
	SkuId         int64        `xorm:"sku_id" json:"skuId,string"`
	Number        int32        `xorm:"number" json:"number"`
	SalePrice     float32      `xorm:"sale_price" json:"salePrice"`
	Price         float32      `xorm:"price" json:"price"`
	Comment       string       `xorm:"comment" json:"comment"`
	Crt           domain.UtcTime    `xorm:"crt" json:"crt"`
	Lut           domain.UtcTime    `xorm:"lut" json:"-"`
	Status        int16        `xorm:"status" json:"-"`
}

func (c *OrderItem) TableName() string {
	return "order_item"
}