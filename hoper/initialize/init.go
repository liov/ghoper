package initialize

import "flag"

func init() {
	cPath := flag.String("conf", "../../config/config.toml", "配置文件地址")
	flag.Parse()
	initializeConfig(*cPath)
	initializeLog()
	initializeDB()
	initializeBoltDB()
	initializeRedis()
	initializeMongo()
	//CacheInit()
}
