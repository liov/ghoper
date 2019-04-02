package main

import (
	"fmt"
	"github.com/kataras/golog"
	"github.com/kataras/pio"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
)

/**
 * @author     ：lbyi
 * @date       ：Created in 2019/4/2
 * @description：
 */

func init() {
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
	golog.SetTimeFormat("2006-01-02 15:04:05")
	golog.Handle(func(l *golog.Log) bool {
		prefix := golog.GetTextForLevel(l.Level, true)
		pc, file, line, _ := runtime.Caller(6)
		message := fmt.Sprintf("%s [%s] %s %s:%d %s",
			prefix, pio.Yellow(l.FormatTime()), l.Message, file, line, pio.Gray(runtime.FuncForPC(pc).Name()))

		if l.NewLine {
			message += "\n"
		}

		print(message)
		return true
	})
}

func main() {

	golog.Warn("Hey, warning here")
	golog.Info("Hey, warning here")
	golog.Debug("Hey, warning here")
	golog.Error("Something went wrong!")

	logrus.Warn("Hey, warning here")
	logrus.Errorf("Something went wrong!")
	logrus.Info("Something went wrong!")
}
