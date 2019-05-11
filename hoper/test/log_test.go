package test

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"testing"
	"time"

	"github.com/kataras/golog"
	"github.com/kataras/pio"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

/**
 * @author     ：lbyi
 * @date       ：Created in 2019/4/2
 * @description：
 */

func Benchmark_Log(b *testing.B) {
	for i := 0; i < b.N; i++ {
		golog.Error("Hey, warning here")
	}
}


func Benchmark_log(b *testing.B) {
	logger := log.New(os.Stdout, "", log.Ltime)
	for i := 0; i < b.N; i++ {
		logger.Print("警告")
	}
}

func Benchmark_golog(b *testing.B) {
	golog.SetTimeFormat("2006/01/02 15:04:05")
	golog.Handle(func(l *golog.Log) bool {
		//应该是深度，好烦，不带format是6(Info)，带是7(Infof)，无法确定
		_, file, line, _ := runtime.Caller(6)

/*		if l.Logger.Printer.IsTerminal {
			l.Message = fmt.Sprintf("[%s] %s:%d %s",
				pio.Red(l.Message), file, line, pio.Gray(runtime.FuncForPC(pc).Name()))
		} else {
			l.Message = fmt.Sprintf("[%s] %s:%d %s",
				l.Message, file, line, runtime.FuncForPC(pc).Name())
		}*/

		if l.Logger.Printer.IsTerminal {
			l.Message = fmt.Sprintf("[%s] %s:%d",
				pio.Red(l.Message), file, line)
		} else {
			l.Message = fmt.Sprintf("[%s] %s:%d",
				l.Message, file, line)
		}

		if l.NewLine {
			l.Logger.Printer.Println(l)
		} else {
			l.Logger.Printer.Print(l)
		}
		return true
	})

	for i := 0; i < b.N; i++ {
		golog.Warn("警告")
	}
}


func Benchmark_logrus(b *testing.B) {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetReportCaller(true)
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logrus.SetOutput(os.Stderr)
	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	customFormatter.FullTimestamp = true
	customFormatter.ForceColors = true
	logrus.SetFormatter(customFormatter)
	// Only log the warning severity or above.
	logrus.SetLevel(logrus.InfoLevel)
	for i := 0; i < b.N; i++ {
		logrus.Warn("警告")
	}
}

func Benchmark_zap(b *testing.B) {
	config := zap.NewProductionConfig()
	config.EncoderConfig = zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	logger, _ := config.Build()
	defer logger.Sync()
	sugar := logger.Sugar()
	for i := 0; i < b.N; i++ {
		sugar.Warnw("警告",
			"attempt", 3,
			"backoff", time.Second)
	}
}
