package main

import (
	"fastService/controller/common/e"
	"fastService/controller/common/gredis"
	"strings"
)

func main() {
	key := strings.Join([]string{
		e.CacheMoment,
		"List",
	}, "_")
	if gredis.Exists(key) {
		gredis.Delete(key)
	}
}
