package cron

import (
	"github.com/gomodule/redigo/redis"
	"fastService/controller"
	"fastService/controller/common"
	"fastService/controller/common/gredis"
	"fastService/initialize"
)

func RedisToDB() {
	conn := initialize.RedisPool.Get()
	defer conn.Close()

	if gredis.Exists(gredis.TopMoments) {
		topData, _ := redis.Strings(conn.Do("LRANGE", gredis.TopMoments, 0, -1))
		for _, mv := range topData {
			if mv != "" {
				var moment controller.Moment
				common.Json.UnmarshalFromString(mv, &moment)
				initialize.DB.Updates(moment)
			}
		}
	}
	data, _ := redis.Strings(conn.Do("LRANGE", gredis.Moments, 0, -1))
	for _, mv := range data {
		if mv != "" {
			var moment controller.Moment
			common.Json.UnmarshalFromString(mv, &moment)
			initialize.DB.Updates(moment)
		}
	}
}
