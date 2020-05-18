package bootstrap

import (
	"fmt"
	"gomall/app/global"
	"os"

	"github.com/sirupsen/logrus"
)

func SetupLog() {
	global.GlobalLog = logrus.New()
	if global.GlobalConfig.Log.App.Formatter == "json" {
		global.GlobalLog.Formatter = new(logrus.JSONFormatter)
	}
	// logrus.TextFormatter

	file, err := os.OpenFile(global.GlobalConfig.Log.App.File, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err == nil {
		global.GlobalLog.Out = file
	} else {
		fmt.Println("open log file error:", err)
	}
}
