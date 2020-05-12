package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"seckill/core/model"
	"seckill/core/model/request"
	resp "seckill/core/model/responce"
	"seckill/core/service"
	"seckill/global"
	"seckill/global/response"
)

// @Tags user
// @Summary 用户注册账号
// @Produce  application/json
// @Param data body model.User true "用户注册接口"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"注册成功"}"
// @Router /base/register [post]
func Register(c *gin.Context) {
	var R request.RegisterStruct
	_ = c.ShouldBindJSON(&R)
	if R.Phone == "" || R.Salt == "" || R.Password == "" {
		global.LOG.Debugf("Invalid params phone is %s password is %s salt is %s", R.Phone, R.Password, R.Salt)
		response.FailWithMessage("无效的参数", c)
		return
	}
	user := &model.User{Phone: R.Phone, Nickname: R.Nickname, Password: R.Password, Salt: R.Salt, Head: R.Head}
	err, userReturn := service.Register(*user)

	if err != nil {
		response.FailWithDetailed(response.ERROR, resp.UserResponse{User: userReturn}, fmt.Sprintf("%v", err), c)
	} else {
		response.OkDetailed(resp.UserResponse{User: userReturn}, "注册成功", c)
	}
}

// @Tags user
// @Summary 用户注册账号
// @Produce  application/json
// @Param data body model.User true "用户登陆接口"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/register [post]
func Login(c *gin.Context) {
	var login request.LoginStruct
	_ = c.ShouldBindJSON(&login)
	if login.Phone == "" || login.Password == "" {
		global.LOG.Debugf("Invalid params phone is %s password is %s", login.Phone, login.Password)
		response.FailWithMessage("无效的参数", c)
		return
	}
	user := &model.User{Phone: login.Phone, Password: login.Password}
	err, u := service.Login(*user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(resp.UserResponse{User: u}, c)
	}

}
