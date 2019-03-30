package initialize

func init() {
	Setup()
	initializeDB()
	initializeBoltDB()
	initializeRedis()
	initializeMongo()
	CacheInit()
}
