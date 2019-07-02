package initialize

import (
	"fmt"
	"hoper/utils"
	"hoper/utils/ulog"
	"time"
)

/**
 * @author     ：lbyi
 * @date       ：Created in 2019/4/2
 * @description：
 */

func initializeLog() {
	var err error
	if Config.Server.Env == ProductionMode{
		ulog.ZapLog()
	}
	ulog.LogFile, err = utils.OpenLogFile(fmt.Sprintf("%s%s.%s",
		Config.Server.LogSaveName,
		time.Now().Format(Config.Server.TimeFormat),
		Config.Server.LogFileExt,
	), utils.GetLogFilePath(Config.Server.RuntimeRootPath,Config.Server.LogSavePath))
	if err != nil {
		ulog.Fatal(err)
	}
}
