package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"seckill/core/model"
	"seckill/core/model/request"
	resp "seckill/core/model/responce"
	"seckill/core/service"
	"seckill/global/response"
)

// @Tags Base
// @Summary 用户注册账号
// @Produce  application/json
// @Param data body model.User true "用户注册接口"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"注册成功"}"
// @Router /base/register [post]
func Register(c *gin.Context) {

	var R request.RegisterStruct
	_ = c.ShouldBindJSON(&R)
	user := &model.User{Nickname: R.Nickname, Password: R.Password, Salt: R.Salt, Head: R.Head}
	err, userReturn := service.Register(*user)

	if err != nil {
		response.FailWithDetailed(response.ERROR, resp.UserResponse{User: userReturn}, fmt.Sprintf("%v", err), c)
	}
	response.OkDetailed(resp.UserResponse{User: userReturn}, "注册成功", c)
}
