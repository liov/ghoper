package logging

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"service/initialize"
	"service/utils"
	"time"
)

type Level int

var (
	F *os.File

	DefaultPrefix      = ""
	DefaultCallerDepth = 2

	logger     *log.Logger
	logPrefix  = ""
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func init() {
	var err error
	F, err = openLogFile(getLogFileName(), getLogFilePath())
	if err != nil {
		log.Fatalln(err)
	}
	logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Println(v)
}

func Info(v ...interface{}) {
	setPrefix(INFO)
	logger.Println(v)
}

func Warn(v ...interface{}) {
	setPrefix(WARNING)
	logger.Println(v)
}

func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Println(v)
}

func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	logger.Fatalln(v)
}

func setPrefix(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}

	logger.SetPrefix(logPrefix)
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

func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		initialize.Config.Server.LogSaveName,
		time.Now().Format(initialize.Config.Server.TimeFormat),
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

	err = utils.IsNotExistMkDir(src)
	if err != nil {
		return nil, fmt.Errorf("文件不存在 src: %s, err: %v", src, err)
	}

	f, err := utils.Open(src+fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("打开失败 :%v", err)
	}

	return f, nil
}
