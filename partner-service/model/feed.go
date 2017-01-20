package model

import (
	"code.isstream.com/stream/domain"
)

type Feed struct {
	Id         int64            `xorm:"id pk" json:"id,string"`
	CustomerId int64            `xorm:"customer_id" json:"customerId,string"`
	CreatorId  int64            `xorm:"creator_id" json:"creatorId,string"`
	Title      string           `xorm:"title" json:"title"`
	Content    string            `xorm:"content" json:"content"`
	ImageIds   domain.Int64Array `xorm:"img_ids" json:"imgIds"`
	Crt        domain.UtcTime    `xorm:"crt timestampz created" json:"crt"`
	Lut        domain.UtcTime    `xorm:"lut timestampz updated" json:"-"`
	Status     int16            `xorm:"status" json:"-"`
}

type FeedRS struct {
	Id         int64            `xorm:"id pk" json:"id,string"`
	CustomerId int64            `xorm:"customer_id" json:"customerId,string"`
	CreatorId  int64            `xorm:"creator_id" json:"creatorId,string"`
	Title      string           `xorm:"title" json:"title"`
	Content    string            `xorm:"content" json:"content"`
	ImageIds   domain.Int64Array `xorm:"img_ids" json:"imgIds"`
	Crt        domain.UtcTime    `xorm:"crt timestampz created" json:"crt"`
	Lut        domain.UtcTime    `xorm:"lut timestampz updated" json:"-"`
	Status     int16            `xorm:"status" json:"-"`

	CreatorName      string `xorm:"name" json:"name"`
	CreatorAvatarUrl string `xorm:"avatar_url" json:"avatarUrl"`
}

func (model *Feed) TableName() string {
	return "feed"
}