package service

import (
	"errors"
	"seckill/core/model"
	"seckill/global"
	"seckill/utils"
)

func Register(user model.User) (err error, u model.User) {

	notRegister := global.BIZ_DB.Where("phone = ?", user.Phone).First(&user).RecordNotFound()

	if notRegister {
		user.Password = utils.MD5V([]byte(user.Password), []byte(user.Salt))
		err = global.BIZ_DB.Create(&user).Error
	} else {
		return errors.New("手机号已注册"), u
	}
	return err,user
}
