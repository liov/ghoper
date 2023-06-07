package cron

import (
	"hoper/utils/ulog"
)

type RedisTo struct {
}

func (RedisTo) Run() {
	ulog.Info("定时任务执行")
}
