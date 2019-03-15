package common

import (
	"github.com/kataras/golog"
	"hoper/initialize"
	"strconv"
	"strings"
)

var kindIndex = map[string]int{
	"moment":  1,
	"article": 2,
}

func CountToRedis(userId uint, refId uint, kind string, operation string) error {
	conn := initialize.RedisPool.Get()
	defer conn.Close()
	conn.Send("MULTI")
	conn.Send("SELECT", kindIndex[kind])
	conn.Send("SADD", strings.Join([]string{"user", strconv.FormatUint(uint64(userId), 10), kind, operation}, "_"), refId)
	conn.Send("INCR", strings.Join([]string{kind, strconv.FormatUint(uint64(refId), 10), operation, "count"}, "_"))
	conn.Send("SELECT", 0)
	_, err := conn.Do("EXEC")
	if err != nil {
		golog.Error("缓存失败", err)
	}
	return nil
}
