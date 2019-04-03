package test

import (
	"fmt"
	"github.com/kataras/golog"
	"hoper/initialize"
	"hoper/utils"
	"reflect"
	"testing"
)

/**
 * @author     ：lbyi
 * @date       ：Created in 2019/4/3
 * @description：
 */

func TestConfig(t *testing.T) {

	conn := initialize.RedisPool.Get()
	defer conn.Close()

	key := "config"
	conn.Send("MULTI")
	conn.Send("SELECT", 12)
	tp := reflect.TypeOf(initialize.Config)
	value := reflect.ValueOf(initialize.Config)
	for i := 0; i < tp.NumField(); i++ {
		// 获取每个成员的结构体字段类型
		fieldType := tp.Field(i)
		for j := 0; j < fieldType.Type.NumField(); j++ {
			f := tp.FieldByIndex([]int{i, j})
			v := value.FieldByIndex([]int{i, j}).Interface()
			conn.Send("HSET", key, f.Name, v)
			fmt.Println(f.Name, v)
		}
	}
	_, err := conn.Do("EXEC")
	if err != nil {
		golog.Error(err)
	}
}

func TestConfig2(t *testing.T) {
	conn := initialize.RedisPool.Get()
	defer conn.Close()

	conn.Send("SELECT", 12)

	config, err := utils.Json.MarshalToString(initialize.Config)
	conn.Do("SET", "config", config)

	if err != nil {
		golog.Error(err)
	}
}
