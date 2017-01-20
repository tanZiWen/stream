package db

import (
	"github.com/go-xorm/xorm"
	"code.isstream.com/stream/setting"
	_"github.com/lib/pq"
	"database/sql"
	log "github.com/Sirupsen/logrus"
	"errors"
)


func InitializePq() error {
	var err error
	if !setting.Initialized {
		err = errors.New("try to initialize postgresql before global setting initialized")
		log.Error(err)
		panic(err)
	}

	Engine, err = xorm.NewEngine(sql.Drivers()[0], config.Rdb.ConnectionString)
	if err != nil {
		log.Error("failed to init postgresql connection ", err)
		return errors.New("connection initializing error")
	}

	Engine.SetMaxOpenConns(config.Rdb.MaxConnections)
	Engine.SetMaxIdleConns(config.Rdb.MaxIdleConnections)
	Engine.ShowSQL()

	return nil
}
