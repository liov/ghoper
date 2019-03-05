package cron

import (
	"github.com/gomodule/redigo/redis"
	"hoper/client/controller"
	"hoper/client/controller/common"
	"hoper/client/controller/common/gredis"
	"hoper/initialize"
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
				initialize.DB.Model(&moment).UpdateColumns(controller.Moment{CollectCount: moment.CollectCount,
					BrowseCount: moment.BrowseCount, CommentCount: moment.CommentCount,
					LikeCount: moment.LikeCount,
				})
			}
		}
	}
	data, _ := redis.Strings(conn.Do("LRANGE", gredis.Moments, 0, -1))
	for _, mv := range data {
		if mv != "" {
			var moment controller.Moment
			common.Json.UnmarshalFromString(mv, &moment)
			initialize.DB.Model(&moment).UpdateColumns(controller.Moment{CollectCount: moment.CollectCount,
				BrowseCount: moment.BrowseCount, CommentCount: moment.CommentCount,
				LikeCount: moment.LikeCount,
			})
		}
	}
}
