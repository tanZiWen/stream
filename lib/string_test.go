package lib

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_Pq(t *testing.T) {
	Convey("customer", t, func() {

		Convey("mask english name", func() {
			Convey("one character", func() {
				name := "m"
				maskedName := MaskName(name)

				t.Log("mask name ", name, maskedName)

				So(maskedName, ShouldNotBeBlank)
			})
			Convey("long", func() {
				name := "Morgan Wu"
				maskedName := MaskName(name)

				t.Log("mask name ", name, maskedName)

				So(maskedName, ShouldNotBeBlank)
			})
			Convey("short", func() {
				name := "ai"
				maskedName := MaskName(name)

				t.Log("mask name ", name, maskedName)

				So(maskedName, ShouldNotBeBlank)
			})
		})

		Convey("mask chinese name", func() {
			Convey("one character", func() {
				name := "吴"
				maskedName := MaskName(name)

				t.Log("mask name ", name, maskedName)

				So(maskedName, ShouldNotBeBlank)
			})
			Convey("two characters", func() {
				name := "吴寰"
				maskedName := MaskName(name)

				t.Log("mask name ", name, maskedName)

				So(maskedName, ShouldNotBeBlank)
			})
			Convey("three characters", func() {
				name := "吴梓修"
				maskedName := MaskName(name)

				t.Log("mask name ", name, maskedName)

				So(maskedName, ShouldNotBeBlank)
			})
			Convey("four characters", func() {
				name := "吴梓修名"
				maskedName := MaskName(name)

				t.Log("mask name ", name, maskedName)

				So(maskedName, ShouldNotBeBlank)
			})
			Convey("five characters", func() {
				name := "吴梓修名字"
				maskedName := MaskName(name)

				t.Log("mask name ", name, maskedName)

				So(maskedName, ShouldNotBeBlank)
			})
		})
	})
}