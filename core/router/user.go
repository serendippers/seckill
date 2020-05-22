package router

import (
	"github.com/gin-gonic/gin"
	"seckill/core/api"
	"seckill/core/middleware"
)

func InitUserRouter(apiGroup *gin.RouterGroup) {
	userRouter := apiGroup.Group("user").Use(middleware.JWTAuth())
	{
		userRouter.POST("change_password", api.ChangePassword)
	}
}
