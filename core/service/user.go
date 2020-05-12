package service

import (
	"errors"
	"seckill/core/model"
	"seckill/global"
	"seckill/utils"
	"time"
)

func Register(user model.User) (err error, u model.User) {
	var userResult model.User
	notRegister := global.BIZ_DB.Where("phone = ?", user.Phone).First(&userResult).RecordNotFound()
	if !notRegister {
		return errors.New("手机号已注册"), u
	}
	user.Id, err = global.IdWorker.NextId()
	if err != nil {
		global.LOG.Errorf("IdWorker create id err:%s\n", err)
		return err, user
	}
	user.Password = utils.MD5V([]byte(user.Password), []byte(user.Salt))
	err = global.BIZ_DB.Create(&user).Error
	return err, user
}

func Login(user model.User) (err error, u model.User) {
	var userResult model.User
	err = global.RO_DB.Where("phone = ?", user.Phone).First(&userResult).Error
	if err != nil {
		global.LOG.Error("login error is ", err)
		return err, u
	}
	password := utils.MD5V([]byte(user.Password), []byte(userResult.Salt))
	notMatch := global.RO_DB.Where("phone = ? and password = ?", user.Phone, password).First(&userResult).RecordNotFound()
	if notMatch {
		return errors.New("密码不正确"), u
	}
	userResult.LastLoginDate = time.Now()
	userResult.LoginCount++
	global.BIZ_DB.Model(&userResult).Updates(model.User{LoginCount: userResult.LoginCount, LastLoginDate: userResult.LastLoginDate})
	return nil, userResult
}
