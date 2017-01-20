package model

import (
	"code.isstream.com/stream/domain"
)

type Sku struct {
	Id          int64            `xorm:"id pk" json:"id,string"`
	Name        string           `xorm:"name" json:"name"`
	SpuId       int64            `xorm:"spu_id" json:"spuId"`
	ItemId      int64            `xorm:"item_id" json:"itemId"`
	PartnerId   int64            `xorm:"partner_id" json:"partnerId"`
	Description string            `xorm:"description" json:"description"`
	SaleStatus  int16            `xorm:"sale_status" json:"saleStatus"`
	StockNumber int              `xorm:"stock_number" json:"stockNumber"`
	PricingType int16            `xorm:"pricing_type" json:"pricingType"`
	MarketPrice float32            `xorm:"market_price" json:"marketPrice"`
	SalePrice   float32            `xorm:"sale_price" json:"salePrice"`
	CoverId     int64            `xorm:"cover_id" json:"coverId"`
	CoverUrl    string            `xorm:"cover_url" json:"coverUrl"`
	PublishTime domain.UtcTime        `xorm:"publish_time timestampz created" json:"publishTime"`
	Crt         domain.UtcTime        `xorm:"crt timestampz created" json:"crt"`
	Lut         domain.UtcTime        `xorm:"lut timestampz updated" json:"-"`
	Status      int16            `xorm:"status" json:"-"`
}

func (model *Sku) TableName() string {
	return "item"
}