package initialize

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"seckill/global"
)

func Mysql() {
	BizMysql()
	RoMysql()
}

func BizMysql() {
	bizConfig := global.CONFIG.BizMySQL
	source := bizConfig.Username + ":" + bizConfig.Password + "@(" + bizConfig.Path + ")/" + bizConfig.Database + "?" + bizConfig.Config
	if bizDB, err := gorm.Open("mysql", source); err != nil {
		global.LOG.Error("BIZ_DB数据库启动异常 source is %s",source, err)
	} else {
		global.BIZ_DB = bizDB
		global.BIZ_DB.DB().SetMaxIdleConns(bizConfig.MaxIdleConns)
		global.BIZ_DB.DB().SetMaxOpenConns(bizConfig.MaxOpenConns)
		global.BIZ_DB.LogMode(bizConfig.LogMode)
		global.LOG.Info("连接BIZ_DB数据库：success")
	}
}

func RoMysql()  {
	roConfig := global.CONFIG.RoMySQL
	source := roConfig.Username + ":" + roConfig.Password + "@(" + roConfig.Path + ")/" + roConfig.Database + "?" + roConfig.Config
	if roDB, err := gorm.Open("mysql", source); err != nil {
		global.LOG.Error("RO_DB数据库启动异常", err)
	} else {
		global.RO_DB = roDB
		global.RO_DB.DB().SetMaxIdleConns(roConfig.MaxIdleConns)
		global.RO_DB.DB().SetMaxOpenConns(roConfig.MaxOpenConns)
		global.RO_DB.LogMode(roConfig.LogMode)
		global.LOG.Info("连接RO_DB数据库：success")
	}
}
