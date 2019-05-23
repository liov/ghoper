package initialize

func init() {
	initializeConfig()
	initializeDB()
	initializeBoltDB()
	initializeRedis()
	initializeMongo()
	//CacheInit()
}
