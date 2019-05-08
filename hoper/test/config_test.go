package test

import (
	"fmt"
	"github.com/kataras/golog"
	"hoper/initialize"
	"hoper/model/crm"
	"hoper/utils"
	"reflect"
	"testing"
	"time"
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

func TestConfig3(t *testing.T) {
	tp := reflect.TypeOf(initialize.Config)
	value := reflect.ValueOf(initialize.Config)
	for i := 0; i < tp.NumField(); i++ {
		// 获取每个成员的结构体字段类型
		fieldType := tp.Field(i)
		d := crm.Dictionary{
			CreatedAt: time.Now(),
			Type:      "config",
			ParentID:  0,
			ParentKey: "",
			Key:       fieldType.Name,
			Value:     fieldType.Name,
			Sequence:  0,
			Status:    1,
		}
		initialize.DB.Create(&d)
		id := d.ID
		for j := 0; j < fieldType.Type.NumField(); j++ {
			f := tp.FieldByIndex([]int{i, j})
			v := value.FieldByIndex([]int{i, j}).Interface()
			d.ParentID = id
			d.ParentKey = d.Key
			d.Key = f.Name
			d.Value = fmt.Sprintf("%v", v)
			d.ID = 0
			initialize.DB.Create(&d)
		}
	}
}
