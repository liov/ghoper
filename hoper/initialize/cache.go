package initialize

import "github.com/bluele/gcache"

var Cache gcache.Cache

func CacheInit() {
	Cache = gcache.New(20).LRU().Build()
}
