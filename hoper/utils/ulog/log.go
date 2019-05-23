package ulog

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"hoper/initialize"
	"hoper/utils"
)

type Logger interface {
	Debug(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Debugf(template string, args ...interface{})
	Errorf(template string, args ...interface{})
	Fatalf(template string, args ...interface{})
	Infof(template string, args ...interface{})
	Warnf(template string, args ...interface{})
}

var (
	LogFile *os.File
	Log     Logger
)

func init() {
	var err error
	initializeLog()
	LogFile, err = openLogFile(getLogFileName(time.Now().Format(initialize.Config.Server.TimeFormat)), getLogFilePath())
	if err != nil {
		Fatal(err)
	}

}

func Debug(v ...interface{}) {
	Log.Debug(v)
}

func Info(v ...interface{}) {
	Log.Info(v)
}

func Warn(v ...interface{}) {
	Log.Warn(v)
}

func Error(v ...interface{}) {
	Log.Error(v)
}

func Fatal(v ...interface{}) {
	Log.Error(v)
}

func Fatalf(template string, v ...interface{}) {
	Log.Error(v)
}

func Debugf(template string, v ...interface{}) {
	Log.Debug(v)
}

func Infof(template string, v ...interface{}) {
	Log.Info(v)
}

func Warnf(template string, v ...interface{}) {
	Log.Warn(v)
}

func Errorf(template string, v ...interface{}) {
	Log.Error(v)
}

func getLogFilePath() string {
	RuntimeRootPath := initialize.Config.Server.RuntimeRootPath
	LogSavePath := initialize.Config.Server.LogSavePath
	if runtime.GOOS == "windows" {
		RuntimeRootPath = RuntimeRootPath + "\\"
		LogSavePath = LogSavePath + "\\"
	} else if runtime.GOOS == "linux" {
		RuntimeRootPath = RuntimeRootPath + "/"
		LogSavePath = LogSavePath + "/"
	}

	return RuntimeRootPath + LogSavePath
}

func getLogFileName(name string) string {
	return fmt.Sprintf("%s%s.%s",
		initialize.Config.Server.LogSaveName,
		name,
		initialize.Config.Server.LogFileExt,
	)
}

func openLogFile(fileName, filePath string) (*os.File, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("os.Getwd err: %v", err)
	}

	src := dir + filePath
	perm := utils.CheckPermission(src)
	if perm == true {
		return nil, fmt.Errorf("权限不足 src: %s", src)
	}

	err = utils.IsNotExistMkdir(src)
	if err != nil {
		return nil, fmt.Errorf("文件不存在 src: %s, err: %v", src, err)
	}

	f, err := utils.Open(src+fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("打开失败 :%v", err)
	}

	return f, nil
}
