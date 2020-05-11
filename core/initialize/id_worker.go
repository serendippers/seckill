package initialize

import (
	"github.com/zheng-ji/goSnowFlake"
	"seckill/global"
)

func CreateIdWorker() {
	if iw, err := goSnowFlake.NewIdWorker(1); err != nil {
		global.LOG.Error("CreateIdWorker goSnowFlake err", err)
	} else {
		global.IdWorker = iw
	}
}
