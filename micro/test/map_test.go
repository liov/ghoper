package test

import (
	"encoding/json"
	"fmt"
	"hoper/model"
	"hoper/utils"
	"log"
	"testing"
)

func TestMap(t *testing.T) {
	l := []interface{}{5, 1, 6}
	utils.Map(func(v interface{}) interface{} {
		i := v.(int)
		return i * i
	}, l)
	fmt.Println(l)
}

type UU struct {
	Name string
	ID   uint
	Sex  string
}

func TestCopy(t *testing.T) {
	u := model.User{Name: "贾一饼", ID: 1, Sex: "男"}
	var uu UU
	if e := utils.CopyProperties(u, &uu); e != nil {
		fmt.Println(e)
	}
	fmt.Println(uu)
}

func TestMap2(t *testing.T) {
	jsonStr := `
    {
        "name":"qqw",
        "id":1,
		"sex":"男"
    }
    `
	u := model.User{}

	//使用 json.Unmarshal(data []byte, v interface{})进行转换,返回 error 信息
	if err := json.Unmarshal([]byte(jsonStr), &u); err != nil {
		log.Fatal(err)
	}
	fmt.Println(u)
}
