package test

import (
	"fmt"
	"hoper/model"
	"hoper/utils"
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

func TestCopyFromBytes(t *testing.T) {
	jsonStr := `
    {
        "name":"贾一饼",
        "id":1,
		"sex":"男"
    }
    `
	u := model.User{}

	//使用 json.Unmarshal(data []byte, v interface{})进行转换,返回 error 信息
	if e := utils.CopyFromBytes(utils.ToBytes(jsonStr), &u); e != nil {
		fmt.Println(e)
	}
	fmt.Println(u)
}
