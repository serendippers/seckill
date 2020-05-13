package api

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"seckill/core/middleware"
	"seckill/core/model"
	"seckill/core/model/request"
	resp "seckill/core/model/responce"
	"seckill/core/service"
	"seckill/global"
	"seckill/global/response"
	"strconv"
	"time"
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
		global.LOG.Debugf("Invalid params phone is %s password is %s salt is %s\n", R.Phone, R.Password, R.Salt)
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
		global.LOG.Debugf("Invalid params phone is %s password is %s\n", login.Phone, login.Password)
		response.FailWithMessage("无效的参数", c)
		return
	}
	user := &model.User{Phone: login.Phone, Password: login.Password}
	err, u := service.Login(*user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	token, err := CreateToken(u)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(resp.UserResponse{User: u, Token: token}, c)
}

func ChangePassword(c *gin.Context) {
	var params request.ChangePassword
	_ = c.ShouldBindJSON(&params)
	if params.Phone == "" || params.Password == "" {
		global.LOG.Debugf("Invalid params phone is %s, Password is %s\n", params.Phone, params.Password)
		response.FailWithMessage("无效的参数", c)
		return
	}
	response.Ok(c)
}

func CreateToken(user model.User) (string, error) {
	j := middleware.NewJWT()

	uuId, _ := global.IdWorker.NextId()
	clams := request.CustomClaims{
		UUId:     uuId,
		NickName: user.Nickname,
		Phone:    user.Phone,
		StandardClaims: jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000),       // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 60*60*24*7), // 过期时间 一周
			Issuer:    "zhangpengpeng",                       //签名的发行者
		},
	}
	token, err := j.CreateToken(clams)
	if err != nil {
		return "", errors.New("获取token失败")
	}
	err = service.SetRedisJWT(strconv.FormatInt(clams.UUId, 10), token)
	return token, err
}
