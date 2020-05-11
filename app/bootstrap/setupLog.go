package bootstrap

import (
	"fmt"
	"os"
	. "gomall/app"
	"github.com/sirupsen/logrus"
)

func SetupLog() {
	GlobalLog = logrus.New()
	if GlobalConfig.Log.App.Formatter == "json" {
		GlobalLog.Formatter = new(logrus.JSONFormatter)
	}
	// logrus.TextFormatter

	file, err := os.OpenFile(GlobalConfig.Log.App.File, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err == nil {
		GlobalLog.Out = file
	} else {
		fmt.Println("open log file error:", err)
	}
}