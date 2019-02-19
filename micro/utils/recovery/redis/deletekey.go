package main

import (
	"hoper/client/controller/common/gredis"
	"strings"
)

func main() {
	key := strings.Join([]string{
		gredis.CacheMoment,
		"List",
	}, "_")
	if gredis.Exists(key) {
		gredis.Delete(key)
	}
}
