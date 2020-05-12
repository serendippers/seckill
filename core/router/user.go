package router

import (
	"github.com/gin-gonic/gin"
	"seckill/core/api"
)

func InitUserRouter(apiGroup *gin.RouterGroup) {
	baseRouter := apiGroup.Group("user")
	{
		baseRouter.GET("show", api.Show)
		baseRouter.POST("register", api.Register)
		baseRouter.POST("login", api.Login)
	}
}
