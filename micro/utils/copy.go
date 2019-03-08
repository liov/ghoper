package utils

import (
	"errors"
	"reflect"
)

func Copy(source interface{}, target interface{}) error {
	if Stype := reflect.TypeOf(source).Kind(); Stype != reflect.Struct || Stype != reflect.Ptr {
		return errors.New("source is not a struct")
	}
	if Ttype := reflect.TypeOf(target).Kind(); Ttype != reflect.Struct || Ttype != reflect.Ptr {
		return errors.New("target is not a struct")
	}
	return nil
}
