package lib

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func Test_Struct(t *testing.T) {
	Convey("Timeex", t, func() {
		Convey("From1970", func() {
			var ct int64 = 1460021404610996

			inttime := From1970(ct)
			secs := inttime.Unix()
			microsecs := inttime.Nanosecond() / 1000

			//rfc3339Str := "2016-04-07T16:09:07:206000+0800"
			//strtime, err := time.Parse(time.RFC3339, rfc3339Str)

			//So(err, ShouldBeNil)

			So(secs, ShouldEqual, 1460021404)
			So(microsecs, ShouldEqual, 610996)
		})

		Convey("MicrosecondsFrom1970", func() {
			t := time.Unix(1460021404, 610996 * 1000)
			microsecs := MicrosecondsFrom1970(t)

			So(microsecs, ShouldEqual, 1460021404610996)
		})
	})
}