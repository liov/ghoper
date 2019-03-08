package test

import (
	"fmt"
	"hoper/model"
	"hoper/utils"
	"reflect"
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

func TestCopy(t *testing.T) {
	u := model.User{}
	Stype := reflect.TypeOf(u).Kind()
	fmt.Println(Stype == reflect.Struct)
}
