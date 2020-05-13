package service

import (
	"fmt"
	"seckill/global"
	"time"
)

func SetRedisJWT(userId, token string) (err error) {
	err = global.REDIS.Set(userId, token, time.Second*60*60*24*7).Err()
	if err!=nil {
		err = fmt.Errorf("SetRedisJWT redis err : %q\n", err)
		global.LOG.Error(err)
	}
	return err
}

func GetRedisJWT(userId string) (string, error) {
	return global.REDIS.Get(userId).Result()
}
