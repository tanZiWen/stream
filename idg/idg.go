package idg

import (
	"github.com/Sirupsen/logrus"
	"code.isstream.com/stream/setting"
)

var (
	worker *IdWorker
	config *Config
	sectionName = "idg"
)

type Config struct {
	DatacenterNumber int `ini:"DATACENTER_NUMBER"`
}

func Id() (int64, error) {
	return worker.NextId()
}

func Ids(num int) ([]int64, error) {
	return worker.NextIds(num)
}

func Initialize() {
	var err error
	logrus.Debug("initialize idg")
	worker, err = NewIdWorker(int64(1), int64(config.DatacenterNumber), twepoch); if err != nil {
		panic(err)
	}
}

func init() {
	config = &Config{}
	setting.AddMapping(setting.SectionMap{SectionName: sectionName, MapTo: config})
}
