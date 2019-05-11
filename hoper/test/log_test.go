package test

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

/**
 * @author     ：lbyi
 * @date       ：Created in 2019/4/2
 * @description：
 */

func Benchmark_golog(b *testing.B) {
	logger := log.New(os.Stdout, "", log.Ltime)
	for i := 0; i < b.N; i++ {
		logger.Print("警告")
	}
}

func Benchmark_hlog(b *testing.B) {
	logger := log.New(os.Stdout, "", log.Ltime)
	for i := 0; i < b.N; i++ {
		logger.Print("警告")
	}
}

func Benchmark_logrus(b *testing.B) {

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
