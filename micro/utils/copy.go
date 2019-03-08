package utils

import (
	"errors"
	"reflect"
)

func Copy(source interface{}, tatget interface{}) error {
	if Stype := reflect.TypeOf(source).Kind(); Stype != reflect.Struct || Stype != reflect.Ptr {
		return errors.New("source is not a struct")
	}
	if Ttype := reflect.TypeOf(tatget).Kind(); Ttype != reflect.Struct || Ttype != reflect.Ptr {
		return errors.New("source is not a struct")
	}
	return nil
}
