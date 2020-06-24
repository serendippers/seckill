package service

import (
	"seckill/core/model"
	"seckill/core/model/request"
	"seckill/global"
)

//创建订单
func CreateOrder(order *request.OrderInfo) (message string, code int) {
	message = "生成订单失败"
	code = 7
	var district model.DeliveryAddr
	isInvalid := global.RO_DB.Where("user_id = ? and district_id = ?", order.UserId, order.DeliveryAddrId).
		First(&district).RecordNotFound()
	if isInvalid {
		global.LOG.Errorf("Invalid deliveryAddr, order is %v", order)
		return message, code
	}
	orderInfoId, err := global.IdWorker.NextId()
	if err != nil {
		global.LOG.Errorf("IdWorker error is %v", err)
		return message, code
	}
	var product model.Product
	var seckillProduct model.SeckillProduct

	//创建
	seckillOrder := model.SeckillOrder{
		UserId:    order.UserId,
		ProductId: order.ProductId,
		OrderId:   orderInfoId,
	}

	//开启一个事务，1，2步操作不会产生锁，所以先执行，3，4会在相应的记录上加锁，
	//本身该服务是个秒杀服务，流量会比较大，所以尽量减少一个事务中加锁的时间
	//
	//1.加seckill_orders表
	//2. 加order_infos 表
	//3. 扣减seckill_products表库存
	//4. 扣减products表库存
	global.BIZ_DB.Begin()
	//添加一条用户订单关系记录
	err = global.BIZ_DB.Create(&seckillOrder).Error
	if err != nil {
		global.LOG.Errorf("create seckillOrder err, error is %v", err)
		global.BIZ_DB.Rollback()
		return message, code
	}

	global.BIZ_DB.Where("id = ?", order.ProductId).First(&product)
	orderInfo := model.OrderInfo{
		Id:             orderInfoId,
		UserId:         order.UserId,
		ProductId:      order.ProductId,
		DeliveryAddrId: order.DeliveryAddrId,
		ProductName:    product.Name,
		ProductCount:   1, //每人只能秒杀一次，每次一件商品
		Status:         0,
		ProductPrice:   product.Price,
		OrderChannel:   1, //暂时默认全部是从web端来的请求
	}

	if product.Stock <= 0 {
		global.LOG.Warningf("Inventory shortage productId is %d  , error is %v", product.Id, err)
		global.BIZ_DB.Rollback()
		return message, code
	}

	err = global.BIZ_DB.Create(&orderInfo).Error
	if err != nil {
		global.LOG.Errorf("create orderInfo err, error is %v", err)
		global.BIZ_DB.Rollback()
		return message, code
	}

	err = global.BIZ_DB.Model(&product).Where("stock > 0").Update("stock - 1").Error
	if err != nil {
		global.LOG.Errorf("products deduct inventory err, productId is %d, error is %v", product.Id, err)
		global.BIZ_DB.Rollback()
		return message, code
	}

	err = global.BIZ_DB.Model(&seckillProduct).Where("stock > 0").Update("stock - 1").Error
	if err != nil {
		global.LOG.Errorf("seckill_products deduct inventory err, productId is %d, error is %v", product.Id, err)
		global.BIZ_DB.Rollback()
		return message, code
	}
	global.BIZ_DB.Commit()
	message = "创建订单成功"
	code = 0

	return message, code
}
