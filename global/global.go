

package global

import (
"seckill/config"
"github.com/go-redis/redis"
"github.com/jinzhu/gorm"
"github.com/op/go-logging"
"github.com/spf13/viper"
)

var (
	BIZ_DB *gorm.DB
	RO_DB  *gorm.DB
	REDIS  *redis.Client
	CONFIG config.Server
	LOG    *logging.Logger
	VIPER  *viper.Viper
)

