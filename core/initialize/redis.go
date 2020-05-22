package initialize

import (
	"github.com/go-redis/redis"
	"seckill/core/model"
	"seckill/global"
	"strconv"
	"time"
)

func Redis() {

	redisConfig := global.CONFIG.Redis

	client := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr,
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	})

	pong, err := client.Ping().Result()

	if err != nil {
		global.LOG.Error(err)
	} else {
		global.LOG.Info("redis connect ping response:", pong)
		global.REDIS = client
	}
}

//项目启动时秒杀商品表中的库存信息缓存进redis
func CacheProduct() {
	var list []model.SeckillProduct
	err := global.RO_DB.Where("end_date > ?", time.Now()).Find(&list).Error
	if err != nil {
		global.LOG.Errorf("CacheProduct fail, error is %v/n", err)
	}
	for _, v := range list {
		if v.Stock > 0  {
			expiration := v.EndDate.Sub(v.StartDate) + time.Minute*10
			key := global.CONFIG.RedisPrefix.SeckillStock + strconv.FormatInt(v.Id, 10)
			err = global.REDIS.Set(key, v.Stock, expiration).Err()
			if err!= nil {
				global.LOG.Error(err)
			}
		}
	}
}
