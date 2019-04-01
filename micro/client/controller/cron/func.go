package cron

import (
	"github.com/gomodule/redigo/redis"
	"hoper/client/controller/cachekey"
	"hoper/initialize"
	"hoper/model/ov"
	"hoper/utils"
	"hoper/utils/gredis"
)

func RedisToDB() {
	conn := initialize.RedisPool.Get()
	defer conn.Close()

	if gredis.Exists(cachekey.TopMoments) {
		topData, _ := redis.Strings(conn.Do("LRANGE", cachekey.TopMoments, 0, -1))
		for _, mv := range topData {
			if mv != "" {
				var moment ov.Moment
				utils.Json.UnmarshalFromString(mv, &moment)
				initialize.DB.Model(&moment).UpdateColumns(ov.Moment{
					ActionCount: ov.ActionCount{
						CollectCount: moment.CollectCount,
						BrowseCount:  moment.BrowseCount, CommentCount: moment.CommentCount,
						LikeCount: moment.LikeCount},
				})
			}
		}
	}
	data, _ := redis.Strings(conn.Do("LRANGE", cachekey.Moments, 0, -1))
	for _, mv := range data {
		if mv != "" {
			var moment ov.Moment
			utils.Json.UnmarshalFromString(mv, &moment)
			initialize.DB.Model(&moment).UpdateColumns(ov.Moment{
				ActionCount: ov.ActionCount{
					CollectCount: moment.CollectCount,
					BrowseCount:  moment.BrowseCount, CommentCount: moment.CommentCount,
					LikeCount: moment.LikeCount},
			})
		}
	}
}
