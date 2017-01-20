package domain

type File struct {
	Id        int64        `xorm:"id pk"`
	Type      int16        `xorm:"type"`
	UserId    int64        `xorm:"user_id"`
	Name      string        `xorm:"name"`
	Ext       string        `xorm:"ext"`
	Size      int            `xorm:"size"`
	OriginUrl string        `xorm:"origin_url"`
	Url       string        `xorm:"url"`
	Meta      map[string]interface{} `xorm:"meta"`
	Crt       MicrosecTime    `xorm:"crt timestampz created" json:"crt"`
	Lut       MicrosecTime    `xorm:"lut timestampz updated" json:"lut"`
	Status    int16          `xorm:"status" json:"status"`
}

func (c *File) TableName() string {
	return "file"
}

type FileRefenence struct {
	Id  int64        `json:"id,string"`
	Url string        `json:"url"`
}
