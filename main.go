package main

import (
	"fmt"
	"github.com/zheng-ji/goSnowFlake"
	"net/http"
	"seckill/core/initialize"
	"seckill/global"
	"seckill/utils"
	"time"
)

func main() {

	iw, _ := goSnowFlake.NewIdWorker(1)
	i,_:= iw.NextId()

	fmt.Printf("%T, %d", i, i)

	//初始化mysql
	initialize.BizMysql()
	//initialize.Redis()
	engine := initialize.Routers()

	initialize.CreateTables()


	defer utils.CrawlerClose()

	s := &http.Server{
		Addr:           "127.0.0.1:8080",
		Handler:        engine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	time.Sleep(10 * time.Microsecond)
	global.LOG.Info("server run success on 8080")


	fmt.Printf("欢迎使用 crawler 默认自动化文档地址:http://%s/swagger/index.html\n", s.Addr)
	global.LOG.Error(s.ListenAndServe())
}
