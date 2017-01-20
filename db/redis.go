package db

import (
	"time"
	"gopkg.in/redis.v2"
	"code.isstream.com/stream/setting"
	log "github.com/Sirupsen/logrus"
	"errors"
)
/*******************EXAMPLE**********************
* Reference: https://godoc.org/gopkg.in/redis.v2
*
************************************************/


var (
	Redis *redis.Client
	//redisConfig *RedisConfig
)

type RedisConfig struct {
	Protocol         string `ini:"PROTOCOL"`
	Address          string    `ini:"ADDRESS"`
	DB               int64    `ini:"DB"`
	Password         string    `ini:"PASSWORD"`
	DialTimeoutSecs  int    `ini:"DIAL_TIMEOUT_SECONDS"`
	ReadTimeoutSecs  int    `ini:"READ_TIMEOUT_SECONDS"`
	WriteTimeoutSecs int    `ini:"WRITE_TIMEOUT_SECONDS"`
	IdleTimeoutMins  int    `ini:"DIAL_TIMEOUT_MINUTES"`
	MaxConnections   int    `ini:"MAX_CONNECTIONS"`
}

func InitializeRedis() error {
	if !setting.Initialized {
		err := errors.New("try to initialize db before global setting initialized")
		log.Error(err)
		panic(err)
	}

	Redis = redis.NewClient(&redis.Options{
		Network: config.Redis.Protocol,
		Addr: config.Redis.Address,
		DB: config.Redis.DB,
		Password: config.Redis.Password,

		DialTimeout:  time.Duration(config.Redis.DialTimeoutSecs) * time.Second,
		ReadTimeout:  time.Duration(config.Redis.ReadTimeoutSecs) * time.Second,
		WriteTimeout: time.Duration(config.Redis.WriteTimeoutSecs) * time.Second,
		IdleTimeout: time.Duration(config.Redis.IdleTimeoutMins) * time.Minute,

		PoolSize:    config.Redis.MaxConnections,
	})

	return nil
}

func init() {
	//redisConfig = &RedisConfig{
	//	Protocol: "tcp",
	//	Address: ":6379",
	//	DialTimeoutSecs: 5,
	//	ReadTimeoutSecs: 1,
	//	WriteTimeoutSecs: 1,
	//	IdleTimeoutMins: 10,
	//	MaxConnections: 50,
	//}
	//setting.AddMapping(setting.SectionMap{SectionName: "database.redis", MapTo: config})
}

