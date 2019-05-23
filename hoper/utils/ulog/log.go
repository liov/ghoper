package ulog

import (
	"fmt"
	"os"
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


func Debug(v ...interface{}) {
	Log.Debug(v...)
}

func Info(v ...interface{}) {
	Log.Info(v...)
}

func Warn(v ...interface{}) {
	Log.Warn(v...)
}

func Error(v ...interface{}) {
	Log.Error(v...)
}

func Fatal(v ...interface{}) {
	Log.Error(v...)
}

func Fatalf(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	Log.Error(msg)
}

func Debugf(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	Log.Debug(msg)
}

func Infof(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	Log.Info(msg)
}

func Warnf(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	Log.Warn(msg)
}

func Errorf(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	Log.Error(msg)
}
