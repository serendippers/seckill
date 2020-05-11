package global

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/op/go-logging"
	"github.com/spf13/viper"
	"github.com/zheng-ji/goSnowFlake"
	"seckill/config"
)

var (
	BIZ_DB   *gorm.DB
	RO_DB    *gorm.DB
	REDIS    *redis.Client
	CONFIG   config.Server
	LOG      *logging.Logger
	VIPER    *viper.Viper
	IdWorker *goSnowFlake.IdWorker
)
