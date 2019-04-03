package initialize

import "time"

func init() {
	initializeConfig()
	initializeLog()
	initializeDB()
	initializeBoltDB()
	initializeRedis()
	initializeMongo()
	//CacheInit()
}

var StartTime = time.Now().Unix()
