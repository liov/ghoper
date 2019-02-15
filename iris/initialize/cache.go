package initialize

import "github.com/bluele/gcache"

var gc gcache.Cache

func CacheInit() {
	gc = gcache.New(20).LRU().Build()
}
