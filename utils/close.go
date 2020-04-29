package utils

import "seckill/global"

func CrawlerClose()  {

	if global.BIZ_DB == nil {
		global.LOG.Info("Biz_DB is null")
		return
	}

	if err := global.BIZ_DB.Close(); err != nil {
		global.LOG.Error("Biz_DB close err ", err)
	}
}
