package initialize

import (
	"flag"
	"fmt"
	"github.com/kataras/golog"
	"hoper/utils/ulog"
	"os"
	"reflect"
	"runtime"
	"time"

	"github.com/jinzhu/configor"
	"hoper/client/controller/credis"
	"hoper/model/crm"
	"hoper/utils"
)

type ServerConfig struct {
	Env          string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PassSalt        string
	TokenMaxAge     int64
	TokenSecret     string
	JwtSecret       string
	PageSize        int8
	RuntimeRootPath string

	UploadDir      string
	UploadPath     string
	UploadMaxSize  int
	UploadAllowExt []string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string

	SiteName string
	Host     string

	MailHost     string
	MailPort     string
	MailUser     string
	MailPassword string

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
		os.Exit(-1)
	}

	if runtime.GOOS == "windows" {
		Config.Server.LuosimaoAPIKey = ""
		Config.Redis.Password = ""
		Config.Server.Env = Debug
	}else {
		flag.StringVar(&Config.Database.Password,"p", Config.Database.Password, "password")
		flag.StringVar(&Config.Server.MailPassword,"mp", Config.Server.MailPassword, "password")
		flag.Parse()
		Config.Redis.Password = Config.Database.Password
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
	conn.Send("SELECT", credis.SysIndex)
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
		ulog.Error(err)
	}
}

func configToRedis2() {
	conn := RedisPool.Get()
	defer conn.Close()

	conn.Send("SELECT", credis.SysIndex)

	config, err := utils.Json.MarshalToString(Config)
	conn.Do("SET", "config", config)

	if err != nil {
		ulog.Error(err)
	}
}

func configToDB() {

	tp := reflect.TypeOf(Config)
	value := reflect.ValueOf(Config)
	for i := 0; i < tp.NumField(); i++ {
		// 获取每个成员的结构体字段类型
		fieldType := tp.Field(i)
		d := crm.Dictionary{
			CreatedAt: time.Now(),
			Type:      "config",
			ParentID:  0,
			ParentKey: "",
			Key:       fieldType.Name,
			Value:     fieldType.Name,
			Sequence:  0,
			Status:    1,
		}
		DB.Create(&d)
		pid := d.ID
		for j := 0; j < fieldType.Type.NumField(); j++ {
			f := tp.FieldByIndex([]int{i, j})
			v := value.FieldByIndex([]int{i, j}).Interface()
			d.ParentID = pid
			d.ParentKey = d.Key
			d.Key = f.Name
			d.Value = fmt.Sprintf("%v", v)
			d.ID = 0
			DB.Create(&d)
		}
	}
}
