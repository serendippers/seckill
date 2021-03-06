package main

import (
	"fmt"
	"net/http"
	"seckill/core/initialize"
	"seckill/global"
	"seckill/utils"
	"time"
)

func main() {

	//初始化mysql
	initialize.Mysql()
	initialize.Redis()
	initialize.InitRabbitMQ()
	//缓存参与秒杀的商品库存
	initialize.CacheProduct()
	engine := initialize.Routers()

	//init goSnowFlake
	initialize.CreateIdWorker()


	initialize.CreateTables()

	defer utils.CrawlerClose()

	s := &http.Server{
		Addr:           "0.0.0.0:8000",
		Handler:        engine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	time.Sleep(10 * time.Microsecond)
	global.LOG.Info("server run success on 8000")


	fmt.Printf("欢迎使用 crawler 默认自动化文档地址:http://%s/swagger/index.html\n", s.Addr)
	global.LOG.Error(s.ListenAndServe())
}
