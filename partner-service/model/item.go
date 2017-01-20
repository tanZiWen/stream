package model

import (
	"code.isstream.com/stream/domain"
)

type Item struct {
	Id           int64            `xorm:"id pk" json:"id,string"`
	Name         string           `xorm:"name" json:"name"`
	SpuId        int64            `xorm:"spu_id" json:"spuId"`
	PartnerId    int64            `xorm:"partner_id" json:"partnerId"`
	CategoryIds  domain.Int64Array            `xorm:"cate_ids" json:"cateIds"`
	Description  string            `xorm:"description" json:"description"`
	ImageIds     domain.Int64Array            `xorm:"img_ids" json:"imgIds"`
	LeadingSkuId int64            `xorm:"leading_sku_id" json:"leadingSkuId"`
	SaleStatus   int16            `xorm:"sale_status" json:"saleStatus"`
	PublishTime  domain.UtcTime        `xorm:"publish_time timestampz created" json:"publishTime"`
	Crt          domain.UtcTime        `xorm:"crt timestampz created" json:"crt"`
	Lut          domain.UtcTime        `xorm:"lut timestampz updated" json:"-"`
	Status       int16            `xorm:"status" json:"-"`
}

func (model *Item) TableName() string {
	return "item"
}