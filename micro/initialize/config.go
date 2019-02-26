package initialize

import (
	"github.com/jinzhu/configor"
	"log"
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
	FileMaxSize    int
	ImageAllowExts []string

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

func Setup() {
	/*Cfg, err := ini.Load("../config/config.ini")
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

	err := configor.New(&configor.Config{Debug: false}).Load(&Config, "../config/config.toml")

	if err != nil {
		log.Fatalf("配置错误: %v", err)
	}

	Config.Server.FileMaxSize = Config.Server.FileMaxSize * 1024 * 1024
	Config.Server.ReadTimeout = Config.Server.ReadTimeout * time.Second
	Config.Server.WriteTimeout = Config.Server.ReadTimeout * time.Second
	Config.Redis.IdleTimeout = Config.Redis.IdleTimeout * time.Second
}
