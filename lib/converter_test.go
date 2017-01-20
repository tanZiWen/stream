package lib

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func Test_DateConverter(t *testing.T) {
	Convey("Given a RFC3339 date string", t, func() {
		var dateString string = "2008-09-08T22:47:31-00:00"

		Convey("The convert should pass", func() {
			parsedDate, err := Str2date(dateString)
			So(err, ShouldBeNil)

			Convey("Value should be equal", func() {
				date := time.Date(2008, 9, 8,22, 47, 31, 0, time.UTC)

				So(parsedDate.UTC().Equal(date), ShouldBeTrue)
			})

			Convey("Default value should be equal", func() {
				defaultDate := time.Date(2008, 9, 8,22, 47, 31, 0, time.UTC)
				d, _ := Str2dateval("", defaultDate)
				So(d.Equal(defaultDate), ShouldBeTrue)
			})
		})
	})
}

func Test_Int64Converter(t *testing.T) {
	Convey("Given a number string", t, func() {
		var str string = "676830858463674369"

		Convey("The convert should pass", func() {
			i, err := Str2int64(str)
			So(err, ShouldBeNil)

			Convey("Value should be equal", func() {
				So(i, ShouldEqual, 676830858463674369)
			})
		})

		Convey("Default value should be equal", func() {
			i, _ := Str2int64val("", 1)
			So(i, ShouldEqual, 1)
		})
	})
}