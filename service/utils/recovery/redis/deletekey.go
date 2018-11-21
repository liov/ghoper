package main

import (
	"service/controller/common/e"
	"service/controller/common/gredis"
	"strings"
)

func main()  {
	key := strings.Join([]string{
		e.CacheMoment,
		"List",
	}, "_")
	if gredis.Exists(key) {
		gredis.Delete(key)
	}
}
