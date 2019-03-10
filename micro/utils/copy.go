package utils

import (
	"errors"
	"reflect"
)

func Copy(source interface{}, target interface{}) error {
	typeOfS := reflect.TypeOf(source)
	typeOfT := reflect.TypeOf(target)

	if typeOfS.Kind() != reflect.Struct && (typeOfS.Kind() == reflect.Ptr && typeOfS.Elem().Kind() != reflect.Struct) && typeOfS.Kind() != reflect.String {
		return errors.New("source is not a struct or string")
	}
	if typeOfS.Kind() == reflect.Ptr {
		typeOfS = reflect.TypeOf(source).Elem()
	}
	if typeOfT.Kind() != reflect.Ptr && typeOfT.Elem().Kind() != reflect.Struct {
		return errors.New("target is not a ptr for struct")
	}
	valueOfS := reflect.ValueOf(source)
	typeOfT = reflect.TypeOf(target).Elem()
	valueOfT := reflect.ValueOf(target).Elem()
	for i := 0; i < typeOfT.NumField(); i++ {
		// 获取每个成员的结构体字段类型
		fieldType := typeOfT.Field(i)
		// 输出成员名和tag
		valueOfT.Field(i).Set(valueOfS.FieldByName(fieldType.Name))
	}

	return nil
}
