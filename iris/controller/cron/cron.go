package cron

import (
	"github.com/gomodule/redigo/redis"
	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
	"service/controller"
	"service/controller/common"
	"service/controller/common/gredis"
	"service/initialize"
)

func StartCron() {

	c := cron.New()
	c.AddFunc("0 0 2 * * *", RedisToDB)
	c.AddJob("0 0 2 * * *", RedisTo{})
	c.Start()
	defer c.Stop()
}

func RedisToDB() {
	conn := initialize.RedisPool.Get()
	defer conn.Close()

	if gredis.Exists(gredis.TopMoments) {
		topData, _ := redis.Strings(conn.Do("LRANGE", gredis.TopMoments, 0, -1))
		for _, mv := range topData {
			if mv != "" {
				var moment controller.Moment
				common.Json.UnmarshalFromString(mv, &moment)
				initialize.DB.Update(&moment)
			}
		}
	}
	data, _ := redis.Strings(conn.Do("LRANGE", gredis.Moments, 0, -1))
	for _, mv := range data {
		if mv != "" {
			var moment controller.Moment
			common.Json.UnmarshalFromString(mv, &moment)
			initialize.DB.Update(&moment)
		}
	}
}

type RedisTo struct {
}

func (RedisTo) Run() {
	logrus.Info("定时任务执行")
}
