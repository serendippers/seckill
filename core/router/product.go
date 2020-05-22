package router

import (
	"github.com/gin-gonic/gin"
	"seckill/core/api"
	"seckill/core/middleware"
)

func InitProductRouter(apiGroup *gin.RouterGroup) {
	productRouter := apiGroup.Group("product").Use(middleware.JWTAuth())
	{
		//productRouter.GET("/:id", api.Product)
		productRouter.GET("/list", api.ProductList)

	}
}
