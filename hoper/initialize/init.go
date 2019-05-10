package initialize

func init() {
	initializeConfig()
	initializeLog()
	initializeDB()
	initializeBoltDB()
	initializeRedis()
	initializeMongo()
	//CacheInit()
}
