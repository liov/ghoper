package test

import (
	"github.com/kataras/golog"
	"hoper/initialize"
	"hoper/model/ov"
	"testing"
)

/**
 * @author     ：lbyi
 * @date       ：Created in 2019/4/1
 * @description：
 */
//ERR wrong number of arguments for 'hmset' command

type User struct {
	ov.User
	ov.KindOwnCount
	Role uint8 `gorm:"type:smallint;default:0" json:"-"` //管理员or用户
}

func TestRedis(t *testing.T) {
	conn := initialize.RedisPool.Get()
	defer conn.Close()

	user := User{
		User: ov.User{
			ID:     1,
			Name:   "贾一饼",
			Status: 1,
		},
		Role: 1,
	}
	data, err := conn.Do("SET", user.ID, user)
	if err != nil {
		golog.Error(err)
	} else {
		golog.Info(data)
	}
}
