package usermodel

import "time"

type User struct {
	Id        int64   `xorm:"id pk" json:"id,string"`
	UserName  string   `xorm:"username" json:"username"`
	Password  string   `xorm:"password" json:"-"`
	Nickname  string   `xorm:"nickname" json:"nickname"`
	Email     string   `xorm:"email" json:"email"`
	Mcc       int16   `xorm:"mcc" json:"-"`
	Mobile    int64   `xorm:"mobile" json:"mobile"`
	//WXOpenid string      `xorm:"wx_openid"`
	AvatarId  string      `xorm:"avatar_id" json:"-"`
	AvatarUrl string     `xorm:"avatar_url" json:"avatarUrl"`
	Language  int `xorm:"language" json:"language" json:"-"`
	Crt       time.Time   `xorm:"crt timestampz created" json:"-"`
	Lut       time.Time   `xorm:"lut timestampz updated" json:"-"`
	Status    int16    `xorm:"status" json:"-"`
}

type UserNickname struct {
	Id           int64   `xorm:"id pk" json:"id,string"`
	Nickname     string   `xorm:"nickname" json:"nickname"`
	Status       int16    `xorm:"status" json:"-"`
}

func (c *User) TableName() string {
	return "app_user"
}