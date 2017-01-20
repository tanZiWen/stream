package db

import (
	"github.com/go-xorm/xorm"
	"code.isstream.com/stream/setting"
	log "github.com/Sirupsen/logrus"
	"errors"
)

var (
	Engine  *xorm.Engine
	config *DatabaseConfig
)

type RdbConfig struct {
	Provider           string `ini:"PROVIDER"`
	ConnectionString   string    `ini:"CONNECTION_STRING"`
	MaxIdleConnections int    `ini:"MAX_IDLE_CONNECTIONS"`
	MaxConnections     int    `ini:"MAX_CONNECTIONS"`
}

type DatabaseConfig struct {
	EnableRdb   bool `ini:"ENABLE_RDB"`
	EnableRedis bool `ini:"ENABLE_REDIS"`

	Redis       *RedisConfig
	Rdb         *RdbConfig
}

func Initialize() error {
	var err error
	if !setting.Initialized {
		err = errors.New("try to initialize db before global setting initialized")
		log.Error(err)
		panic(err)
	}

	if config.EnableRdb {
		log.Debug("rational database enabled")
		if config.Rdb.Provider == "postgresql" {
			err = InitializePq(); if err != nil {
				return err
			}
		} else {
			return errors.New("unsupport database provider, only support postgresql right now.")
		}
	}

	if config.EnableRedis {
		log.Debug("redis database enabled")
		err = InitializeRedis(); if err != nil {
			return err
		}
	}

	return nil
}

func init() {
	rdbConfig := &RdbConfig{
		Provider: "postgresql",
		MaxIdleConnections: 10,
		MaxConnections: 50,
	}

	redisConfig := &RedisConfig{
		Protocol: "tcp",
		Address: ":6379",
		DialTimeoutSecs: 5,
		ReadTimeoutSecs: 1,
		WriteTimeoutSecs: 1,
		IdleTimeoutMins: 10,
		MaxConnections: 50,
	}

	config = &DatabaseConfig{Rdb: rdbConfig, Redis: redisConfig}
	setting.AddMapping(setting.SectionMap{SectionName: "database", MapTo: config})
	setting.AddMapping(setting.SectionMap{SectionName: "database.rdb", MapTo: config.Rdb})
	setting.AddMapping(setting.SectionMap{SectionName: "database.redis", MapTo: config.Redis})
}