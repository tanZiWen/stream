package setting

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

type CustomConfig struct {
	TestKey int `ini:"TEST_KEY"`
}

func Test_Setting(t *testing.T) {
	Convey("Load", t, func() {

		Convey("Default", func() {

			err := Initialize()

			So(err, ShouldBeNil)

			t.Log("page config: ", Page)
			So(Page.Size, ShouldBeGreaterThan, 0)
		})

		Convey("Custom", func() {
			Config = nil
			configLoaded = false

			customConfig := &CustomConfig{}
			AddFile("conf/custom.ini")
			AddMapping(SectionMap{SectionName: "custom", MapTo: customConfig})

			err := Initialize()

			So(err, ShouldBeNil)

			t.Log("custom config: ", customConfig)
			So(customConfig.TestKey, ShouldEqual, 1)
		})
	})
}