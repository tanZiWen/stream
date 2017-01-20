package auth

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"code.isstream.com/stream/auth/setting"
	globalsetting "code.isstream.com/stream/setting"
	"time"
)

func Test_Struct(t *testing.T) {
	Convey("Auth", t, func() {
		Convey("Init", func() {
			globalsetting.Initialize()
			So(setting.Config.Secret, ShouldNotBeNil)

			t.Log(setting.Config)
			t.Log("hour:", int64(time.Hour))
		})
	})
}
