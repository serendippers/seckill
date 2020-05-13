package router

import (
	"github.com/gin-gonic/gin"
	"seckill/core/api"
)

func InitBaseRouter(apiGroup *gin.RouterGroup) {
	baseRouter := apiGroup.Group("base")
	{
		baseRouter.GET("show", api.Show)
		baseRouter.POST("register", api.Register)
		baseRouter.POST("login", api.Login)
	}
}
