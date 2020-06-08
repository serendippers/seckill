package global

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/op/go-logging"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
	"github.com/zheng-ji/goSnowFlake"
	"seckill/resources"
)

type RedisLua struct {
	Path string
	Sha  string
}

var (
	BIZ_DB   *gorm.DB
	RO_DB    *gorm.DB
	REDIS    *redis.Client
	CONFIG   resources.Server
	LOG      *logging.Logger
	VIPER    *viper.Viper
	IdWorker *goSnowFlake.IdWorker
	MQ       *amqp.Connection
	LuaMap   *map[string]*RedisLua
)
