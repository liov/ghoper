package controller

import (
	"fastService/controller/common"
	"github.com/sirupsen/logrus"
	"os"
)

func init() {
	//github.com/gin-contrib/sessions 获取User必须gob注册
	//必须encoding/gob编码解码进行注册
	//gob.Register(&User{})
	// Log as JSON instead of the default ASCII formatter.
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logrus.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	logrus.SetLevel(logrus.InfoLevel)
}

var jsons = common.Json
