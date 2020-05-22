package api

import (
	"github.com/gin-gonic/gin"
	"seckill/core/model/request"
	"seckill/core/model/responce"
	"seckill/core/service"
	"seckill/global/response"
)

func ProductList(c *gin.Context)  {
	var pageInfo request.PageInfo
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.ProductList(&pageInfo)
	if err!= nil {
		response.FailWithMessage("get product list fail: %v", c)
	}
	response.OkWithData(responce.PageResult{
		List: list,
		Total: total,
		Page: pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, c)
}

func SeckillProductList(c *gin.Context)  {
	var pageInfo request.PageInfo
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.SeckillProductList(&pageInfo)
	if err!= nil {
		response.FailWithMessage("get seckill_product list fail: %v", c)
	}
	response.OkWithData(responce.PageResult{
		List: list,
		Total: total,
		Page: pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, c)
}

func Seckill(c *gin.Context) {


}
