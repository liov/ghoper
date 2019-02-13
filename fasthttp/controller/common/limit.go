package common

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"fastService/initialize"
)

func Limit(minuteLimit string, minuteLimitCount int64, dayLimit string, dayLimitCount int64, userID uint) string {

	RedisConn := initialize.RedisPool.Get()
	defer RedisConn.Close()

	minuteKey := minuteLimit + fmt.Sprintf("%d", userID)
	minuteCount, minuteErr := redis.Int64(RedisConn.Do("GET", minuteKey))

	if minuteErr == nil && minuteCount >= minuteLimitCount {
		return "您的操作过于频繁，请先休息一会儿。"
	}

	minuteRemainingTime, _ := redis.Int64(RedisConn.Do("TTL", minuteKey))
	if minuteRemainingTime < 0 || minuteRemainingTime > 60 {
		minuteRemainingTime = 60
	}

	if _, err := RedisConn.Do("SET", minuteKey, minuteCount+1, "EX", minuteRemainingTime); err != nil {
		fmt.Println("redis set failed:", err)
		return "内部错误"
	}

	dayKey := dayLimit + fmt.Sprintf("%d", userID)
	dayCount, dayErr := redis.Int64(RedisConn.Do("GET", dayKey))
	if dayErr == nil && dayCount >= dayLimitCount {
		return "您今天的操作过于频繁，请先休息一会儿。"
	}

	dayRemainingTime, _ := redis.Int64(RedisConn.Do("TTL", dayKey))
	secondsOfDay := int64(24 * 60 * 60)
	if dayRemainingTime < 0 || dayRemainingTime > secondsOfDay {
		dayRemainingTime = secondsOfDay
	}

	if _, err := RedisConn.Do("SET", dayKey, dayCount+1, "EX", dayRemainingTime); err != nil {
		fmt.Println("redis set failed:", err)
		return "内部错误"
	}
	return ""
}
