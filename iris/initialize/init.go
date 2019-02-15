package initialize

func init() {
	Setup()
	initializeDB()
	initializeMongo()
	initializeRedis()
	CacheInit()
}
