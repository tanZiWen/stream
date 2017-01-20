package db

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"code.isstream.com/stream/setting"
)

func Test_Pq(t *testing.T) {
	Convey("db", t, func() {

		Convey("initialize", func() {

			setting.AddFile("conf/test_config.ini")
			err := setting.Initialize()
			So(err, ShouldBeNil)

			t.Log("database config ", config)
			t.Log("redis config ", config.Redis)
			err = Initialize()
			So(err, ShouldBeNil)
			if err != nil {
				t.Error("initialize database error: ", err)
			}
		})
	})
}