package cron

import "github.com/sirupsen/logrus"

type RedisTo struct {
}

func (RedisTo) Run() {
	logrus.Info("定时任务执行")
}
