package service

import (
	"encoding/json"
	"seckill/core/model"
	"seckill/core/model/request"
	"seckill/core/producer"
	"seckill/global"
	"strconv"
	"sync"
)

//内存标记，商品是否被全部秒杀，减少redis访问
//TODO 这种办法感觉很low
var productCache sync.Map

func ProductList(pageInfo *request.PageInfo) (err error, list interface{}, total int) {
	limit := pageInfo.PageSize
	offset := (pageInfo.Page - 1) * pageInfo.PageSize
	db := global.RO_DB
	var productList []model.Product
	err = db.Find(&productList).Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&productList).Error
	if err != nil {
		global.LOG.Errorf("ProductList err: %v\n", err)
	}
	return err, productList, total
}

func SeckillProductList(pageInfo *request.PageInfo) (err error, list interface{}, total int) {
	limit := pageInfo.PageSize
	offset := (pageInfo.Page - 1) * pageInfo.PageSize
	var seckillList []model.SeckillProduct
	err = global.RO_DB.Find(&seckillList).Count(&total).Error
	err = global.RO_DB.Limit(limit).Offset(offset).Find(&seckillList).Error
	if err != nil {
		global.LOG.Errorf("SeckillProductList err: %v\n", err)
	}
	return err, seckillList, total
}

//暂时返回是否成功
func Seckill(info *request.OrderInfo) (message string, ok bool) {
	message = "秒杀失败"
	ok = false
	if _, ok = productCache.Load(info.ProductId); ok {
		global.LOG.Warning("Seckill over\n")
		message = "商品已经秒杀完毕"
		return message, ok
	}
	key := global.CONFIG.RedisPrefix.SeckillStock + strconv.FormatInt(info.ProductId, 10)
	stock, err := global.REDIS.Decr(key).Result()
	if err != nil {
		global.LOG.Errorf("Seckill redis Decr err: %v\n", err)
		return message, ok
	}

	if stock < 0 {
		global.LOG.Warning("Seckill over\n")
		message = "商品已经秒杀完毕"
		productCache.Store(info.ProductId, true)
		return message, ok
	}
	msgJson ,_ := json.Marshal(info)
	producer.ORDER_PRODUCER.SendMessage(msgJson)
	ok = true
	message = "已进入秒杀队列"
	return message, ok
}
