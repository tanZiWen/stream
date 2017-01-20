package domain

import (
	. "github.com/smartystreets/goconvey/convey"
	//log "github.com/Sirupsen/logrus"
	"testing"
	"time"
)

func Test_Time(t *testing.T) {
	Convey("Time", t, func() {
		Convey("utc", func() {
			t1 := time.Now()
			ut := UtcTime(t1)
			utUnmarshal := &UtcTime{}
			//utc := 1478681019230

			bts := &[]byte{}

			var err error
			*bts, err = ut.MarshalJSON()

			t.Log("utc time: ", t1.UnixNano())
			t.Log("utc string: ", string(*bts))

			//utctime := &UtcTime{}
			//err := utctime.UnmarshalJSON(bts)

			So(err, ShouldBeNil)

			//buf := new(bytes.Buffer)
			//err = binary.Write(buf, binary.LittleEndian, t1.UnixNano() / 1000)
			//So(err, ShouldBeNil)

			//*bts = buf.Bytes()
			//*bts = []byte(strconv.FormatInt(t1.UnixNano() / 1000, 10))
			//t.Log("byte array: ", string(*bts))

			var jsonStr string = "\"" + string(*bts) + "\""
			//var newbts []byte = []
			err = utUnmarshal.UnmarshalJSON([]byte(jsonStr))
			So(err, ShouldBeNil)

			t.Log("utc time after unmarshal: ", time.Time(*utUnmarshal).UnixNano(), t1.Unix())
			//
			//t.Log(setting.Config)
			//t.Log("hour:", int64(time.Hour))
		})
	})
}
