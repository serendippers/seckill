package initialize

import (
	"seckill/core/model"
	"seckill/global"
)

func CreateTables() {
	db := global.BIZ_DB
	db.AutoMigrate(
		model.User{},
		model.Product{},
		model.OrderInfo{},
		model.SeckillOrder{},
		model.SeckillProduct{},
		model.DeliveryAddr{},
	)
	global.LOG.Debug("register table success")
}
