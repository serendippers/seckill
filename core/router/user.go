package router

import (
	"github.com/gin-gonic/gin"
	"seckill/core/api"
	"seckill/core/middleware"
)

func InitUserRouter(apiGroup *gin.RouterGroup) {
	baseRouter := apiGroup.Group("user").Use(middleware.JWTAuth())
	{
		baseRouter.POST("change_password", api.ChangePassword)
	}
}
