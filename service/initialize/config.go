package initialize

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)


type ServerConfig struct {
	RunMode string
	Env string
	HttpPort string
	ReadTimeout time.Duration
	WriteTimeout time.Duration

	PassSalt string
	TokenMaxAge time.Duration
	TokenSecret string
	JwtSecret string
	PageSize int
	RuntimeRootPath string

	ImagePrefixUrl string
	UploadImgDir string
	ImagePath string
	ImgHost string
	ImageMaxSize int
	ImageAllowExts []string

	LogSavePath string
	LogSaveName string
	LogFileExt string
	TimeFormat string

	SiteName string
	Host string

	MailHost string
	MailPort int
	MailUser string
	MailPassWord string
	MailFrom string

	LuosimaoVerifyURL string
	LuosimaoAPIKey string

	QrCodeSavePath string	//二维码保存路径
	PrefixUrl	string
	FontSavePath string //字体保存路径

	CrawlerName string  //爬虫
}

var ServerSettings = &ServerConfig{}

type DatabaseConfig struct {
	Type string
	User string
	Password string
	Host string
	Charset string
	Database string
	TablePrefix string
	MaxIdleConns int
	MaxOpenConns int
	Port int
}

var DatabaseSettings = &DatabaseConfig{}

type  RedisConfig struct {
	Host string
	Port int
	Password string
	MaxIdle int
	MaxActive int
	IdleTimeout time.Duration
}

var RedisSettings =&RedisConfig{}

type MongoConfig struct {
	URL string
	Database string
}

var MongoSettings =&MongoConfig{}

func Setup() {
	Cfg, err := ini.Load("../config/config.ini")
	if err != nil {
		log.Fatalf("找不到文件 'website/config/config.ini': %v", err)
	}

	err = Cfg.Section("server").MapTo(ServerSettings)
	if err != nil {
		log.Fatalf("Cfg.MapTo 服务器设置错误: %v", err)
	}

	ServerSettings.ImageMaxSize = ServerSettings.ImageMaxSize * 1024 * 1024
	ServerSettings.ReadTimeout = ServerSettings.ReadTimeout * time.Second
	ServerSettings.WriteTimeout = ServerSettings.ReadTimeout * time.Second
	ServerSettings.TokenMaxAge = ServerSettings.TokenMaxAge *time.Second

	err = Cfg.Section("database").MapTo(DatabaseSettings)
	if err != nil {
		log.Fatalf("Cfg.MapTo 数据库设置错误: %v", err)
	}

	err = Cfg.Section("redis").MapTo(RedisSettings)
	RedisSettings.IdleTimeout = RedisSettings.IdleTimeout * time.Second

	err =Cfg.Section("mongodb").MapTo(MongoSettings)
}

