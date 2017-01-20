package setting

import (
	"time"
	"code.isstream.com/stream/setting"
)

var Config *JwtConfig
var sectionName string = "auth"

type JwtConfig struct {
	Realm        string `ini:"REALM"`
	SignedMethod string `ini:"SIGNED_METHOD"`
	Secret       string `ini:"SECRET"`
	Timeout      int `ini:"TIMEOUT"`
	MaxRefresh   int `ini:"MAX_REFRESH"`
}

func init() {
	Config = &JwtConfig{Realm: "jwt realm", Timeout: int(time.Hour) * 24 * 30 * 12, MaxRefresh: int(time.Hour)}
	setting.AddMapping(setting.SectionMap{SectionName: sectionName, MapTo: Config})
}
