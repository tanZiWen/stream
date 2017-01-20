package lib

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_Mobile(t *testing.T) {
	Convey("Fill mobile number", t, func() {
		prefix := "138"

		suffix := 21

		number := FillNumber(prefix, suffix)

		So(number, ShouldEqual, "13800000021")
	})

	Convey("Range Capability", t, func() {
		prefix := "138"

		capacity, err := RangeCapacity(prefix)

		So(err, ShouldBeNil)
		So(capacity, ShouldEqual, 99999999)
	})
}