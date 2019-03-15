package common

import (
	"errors"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/prometheus/common/log"
	"hoper/initialize"
)

func Limit(minuteLimit string, minuteLimitCount int64, dayLimit string, dayLimitCount int64, userID uint) error {

	RedisConn := initialize.RedisPool.Get()
	defer RedisConn.Close()

	minuteKey := minuteLimit + fmt.Sprintf("%d", userID)
	minuteCount, minuteErr := redis.Int64(RedisConn.Do("GET", minuteKey))

	if minuteErr == nil && minuteCount >= minuteLimitCount {
		return errors.New("您的操作过于频繁，请先休息一会儿。")
	}

	minuteRemainingTime, _ := redis.Int64(RedisConn.Do("TTL", minuteKey))
	if minuteRemainingTime < 0 || minuteRemainingTime > 60 {
		minuteRemainingTime = 60
	}

	if _, err := RedisConn.Do("SET", minuteKey, minuteCount+1, "EX", minuteRemainingTime); err != nil {
		log.Error("redis set failed:", err)
		return errors.New("内部错误")
	}

	dayKey := dayLimit + fmt.Sprintf("%d", userID)
	dayCount, dayErr := redis.Int64(RedisConn.Do("GET", dayKey))
	if dayErr == nil && dayCount >= dayLimitCount {
		return errors.New("您今天的操作过于频繁，请先休息一会儿。")
	}

	dayRemainingTime, _ := redis.Int64(RedisConn.Do("TTL", dayKey))
	secondsOfDay := int64(24 * 60 * 60)
	if dayRemainingTime < 0 || dayRemainingTime > secondsOfDay {
		dayRemainingTime = secondsOfDay
	}

	if _, err := RedisConn.Do("SET", dayKey, dayCount+1, "EX", dayRemainingTime); err != nil {
		log.Error("redis set failed:", err)
		return errors.New("内部错误")
	}
	return nil
}
