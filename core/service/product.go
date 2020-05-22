package service

import (
	"seckill/core/model"
	"seckill/core/model/request"
	"seckill/global"
)


func ProductList(pageInfo *request.PageInfo) (err error, list interface{}, total int) {
	limit := pageInfo.PageSize
	offset := (pageInfo.Page -1) * pageInfo.PageSize
	db := global.RO_DB
	var productList []model.Product
	err = db.Find(&productList).Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&productList).Error
	if err!= nil {
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
	if err!= nil {
		global.LOG.Errorf("SeckillProductList err: %v\n", err)
	}
	return err, seckillList, total

}
