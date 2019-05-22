package initialize

import (
	"fmt"
	"github.com/kataras/golog"
	"github.com/kataras/pio"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"runtime"
)

/**
 * @author     ：lbyi
 * @date       ：Created in 2019/4/2
 * @description：
 */

type Logger interface {
	Debug(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
}

var Log Logger

func initializeLog() {
	// Log as JSON instead of the default ASCII formatter.

	//logrus.SetFormatter(&logrus.JSONFormatter{})

	/*	logrus.SetReportCaller(true)
		// Output to stdout instead of the default stderr
		// Can be any io.Writer, see below for File example
		logrus.SetOutput(os.Stderr)
		formatter :=&logrus.TextFormatter{
			TimestampFormat:"2006-01-02 15:04:05",
			FullTimestamp:true,
			ForceColors:true,
		}
		logrus.SetFormatter(formatter)
		// Only log the warning severity or above.
		logrus.SetLevel(logrus.InfoLevel)*/
	golog.SetTimeFormat("2006/01/02 15:04:05")
	golog.Handle(func(l *golog.Log) bool {
		//应该是深度，好烦，不带format是6，带是7，无法确定
		pc, file, line, _ := runtime.Caller(6)

		if l.Logger.Printer.IsTerminal {
			l.Message = fmt.Sprintf("[%s] %s:%d %s",
				pio.Red(l.Message), file, line, pio.Gray(runtime.FuncForPC(pc).Name()))
		} else {
			l.Message = fmt.Sprintf("[%s] %s:%d %s",
				l.Message, file, line, runtime.FuncForPC(pc).Name())
		}

		if l.NewLine {
			l.Logger.Printer.Println(l)
		} else {
			l.Logger.Printer.Print(l)
		}
		return true
	})
	if Config.Server.Env == "debug"{
		Log = golog.Default
	}else {
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
		Log = logger.Sugar()
	}

}
