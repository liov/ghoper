package cron

import (
	"strconv"
	"sync"
	"time"

	"github.com/gomodule/redigo/redis"
	"hoper/client/controller"
	"hoper/client/controller/credis"
	"hoper/initialize"
	"hoper/model"
)

//这样设计的话，锁的设计就多余了
var StartTime = struct {
	Time int64
	L    sync.RWMutex
}{Time: time.Now().Unix()}

func UserRedisToDB() error {
	conn := initialize.RedisPool.Get()
	defer conn.Close()

	StartTime.L.Lock()
	StartTime.Time = time.Now().Unix() - 3600
	conn.Send("SELECT", credis.CronIndex)
	ids, err := redis.Int64s(conn.Do("ZRANGEBYSCORE", model.LoginUser+"ActiveTime", "-inf", StartTime.Time))
	if err != nil {
		return err
	}
	conn.Do("ZREMRANGEBYSCORE", model.LoginUser+"ActiveTime", "-inf", StartTime.Time)
	StartTime.L.Unlock()

	for id := range ids {
		for i := 0; i <= 4; i++ {
			ua := controller.GetRedisAction(strconv.Itoa(id), int8(i))
			for sid := range ua.Approve {
				err = initialize.DB.Exec("INSERT INTO " + "moment_approve_user VALUES (" +
					strconv.Itoa(sid) + "," + strconv.Itoa(id) + ")").Error
			}
			for sid := range ua.Like {
				err = initialize.DB.Exec("INSERT INTO " + "moment_like_user VALUES (" +
					strconv.Itoa(sid) + "," + strconv.Itoa(id) + ")").Error
			}
			for sid := range ua.Collect {
				err = initialize.DB.Exec("INSERT INTO " + "moment_collect_user VALUES (" +
					strconv.Itoa(sid) + "," + strconv.Itoa(id) + ")").Error
			}
			for sid := range ua.Comment {
				err = initialize.DB.Exec("INSERT INTO " + "moment_comment_user VALUES (" +
					strconv.Itoa(sid) + "," + strconv.Itoa(id) + ")").Error
			}
			/*			for sid:=range ua.Browse{
						err = initialize.DB.Exec("INSERT INTO " + "moment_browse_user VALUES (" +
							strconv.Itoa(sid) + "," + strconv.Itoa(id) + ")").Error
					}*/
		}
	}

	if err != nil {
		return err
	}
	return nil
}
