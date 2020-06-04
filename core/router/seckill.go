package router

import (
	"github.com/gin-gonic/gin"
	"seckill/core/api"
	"seckill/core/middleware"
)

func InitSeckillRouter(apiGroup *gin.RouterGroup) {
	productRouter := apiGroup.Group("seckill").Use(middleware.JWTAuth())
	{
		productRouter.GET("/list", api.SeckillProductList)
		productRouter.POST("/do-seckill", api.Seckill)

	}
}
