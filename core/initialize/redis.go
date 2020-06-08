package initialize

import (
	"github.com/go-redis/redis"
	"io/ioutil"
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

	loadLuaScript()
}

//项目启动时秒杀商品表中的库存信息缓存进redis
func CacheProduct() {
	var list []model.SeckillProduct
	err := global.RO_DB.Where("end_date > ?", time.Now()).Find(&list).Error
	if err != nil {
		global.LOG.Errorf("CacheProduct fail, error is %v/n", err)
	}
	for _, v := range list {
		if v.Stock > 0 {
			expiration := v.EndDate.Sub(v.StartDate) + time.Minute*10
			key := global.CONFIG.RedisPrefix.SeckillStock + strconv.FormatInt(v.Id, 10)
			err = global.REDIS.Set(key, v.Stock, expiration).Err()
			if err != nil {
				global.LOG.Error(err)
			}
		}
	}
}

//将lua脚本缓存至redis中
func loadLuaScript()  {
	luaMap := make(map[string]*global.RedisLua)

	//配置lua脚本的路径
	//秒杀lua脚本
	luaMap["seckill"] = &global.RedisLua{Path: "resources/lua/seckill.lua"}
	//限流lua脚本
	luaMap["limiter"] = &global.RedisLua{Path: "resources/lua/limiter.lua"}
	for k,v := range luaMap {
		date, err := ioutil.ReadFile(v.Path)
		if err != nil {
			global.LOG.Errorf("Unable to read Lua file content, error is %v\n", err)
			panic("读取Lua脚本失败")
		}
		script := redis.NewScript(string(date))
		if result, err := script.Load(global.REDIS).Result(); err != nil {
			panic("缓存脚本到Redis失败")
		} else {
			v.Sha = result
			global.LOG.Infof("%s 对应的sha: %s\n", k, result)
		}
	}
	global.LuaMap = &luaMap
	global.LOG.Info("LoadLuaScript success")
}

