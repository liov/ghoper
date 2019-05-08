package cron

import "github.com/kataras/golog"

type RedisTo struct {
}

func (RedisTo) Run() {
	golog.Info("定时任务执行")
}
