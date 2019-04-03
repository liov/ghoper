package initialize

import (
	"github.com/jinzhu/configor"
	"github.com/kataras/golog"
	"hoper/utils"
	"reflect"
	"time"
)

type ServerConfig struct {
	Env          string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PassSalt        string
	TokenMaxAge     int
	TokenSecret     string
	JwtSecret       string
	PageSize        int
	RuntimeRootPath string

	ImagePrefixUrl string
	UploadDir      string
	UploadPath     string
	UploadHost     string
	UploadMaxSize  int
	UploadAllowExt []string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string

	SiteName string
	Host     string

	MailHost     string
	MailPort     int
	MailUser     string
	MailPassword string
	MailFrom     string

	LuosimaoVerifyURL string
	LuosimaoAPIKey    string

	QrCodeSavePath string //二维码保存路径
	PrefixUrl      string
	FontSavePath   string //字体保存路径

	CrawlerName string //爬虫
}

type DatabaseConfig struct {
	Type         string
	User         string
	Password     string
	Host         string
	Charset      string
	Database     string
	TablePrefix  string
	MaxIdleConns int
	MaxOpenConns int
	Port         int
}

type RedisConfig struct {
	Host        string
	Port        int
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

type MongoConfig struct {
	URL      string
	Database string
}

/*var ServerSettings = &ServerConfig{}
var DatabaseSettings = &DatabaseConfig{}
var RedisSettings = &RedisConfig{}
var MongoSettings = &MongoConfig{}*/

var Config = struct {
	Server   ServerConfig
	Database DatabaseConfig
	Redis    RedisConfig
	Mongo    MongoConfig
}{}

type duration struct {
	time.Duration
}

func (d *duration) UnmarshalText(text []byte) error {
	var err error
	d.Duration, err = time.ParseDuration(string(text))
	return err
}

func initializeConfig() {
	/*Cfg, err := ini.Load("../../config/config.ini")
	if err != nil {
		log.Fatalf("找不到文件 'website/config/config.ini': %v", err)
	}

	err = Cfg.Section("server").MapTo(ServerSettings)
	if err != nil {
		log.Fatalf("Cfg.MapTo 服务器设置错误: %v", err)
	}



	err = Cfg.Section("database").MapTo(DatabaseSettings)
	if err != nil {
		log.Fatalf("Cfg.MapTo 数据库设置错误: %v", err)
	}

	err = Cfg.Section("redis").MapTo(RedisSettings)


	err = Cfg.Section("mongodb").MapTo(MongoSettings)*/

	err := configor.New(&configor.Config{Debug: false}).Load(&Config, "../../config/config.toml")

	if err != nil {
		golog.Error("配置错误: %v", err)
	}

	Config.Server.UploadMaxSize = Config.Server.UploadMaxSize * 1024 * 1024
	Config.Server.ReadTimeout = Config.Server.ReadTimeout * time.Second
	Config.Server.WriteTimeout = Config.Server.WriteTimeout * time.Second
	Config.Redis.IdleTimeout = Config.Redis.IdleTimeout * time.Second
}

func configToRedis() {
	conn := RedisPool.Get()
	defer conn.Close()

	key := "config"
	conn.Send("MULTI")
	conn.Send("SELECT", 12)
	tp := reflect.TypeOf(Config)
	value := reflect.ValueOf(Config)
	for i := 0; i < tp.NumField(); i++ {
		// 获取每个成员的结构体字段类型
		fieldType := tp.Field(i)
		for j := 0; j < fieldType.Type.NumField(); j++ {
			f := tp.FieldByIndex([]int{i, j})
			v := value.FieldByIndex([]int{i, j}).Interface()
			conn.Send("HSET", key, f.Name, v)
		}
	}
	_, err := conn.Do("EXEC")
	if err != nil {
		golog.Error(err)
	}
}

func configToRedis2() {
	conn := RedisPool.Get()
	defer conn.Close()

	conn.Send("SELECT", 12)

	config, err := utils.Json.MarshalToString(Config)
	conn.Do("SET", "config", config)

	if err != nil {
		golog.Error(err)
	}
}
