package main

import (
	"os"
	"fmt"
	"flag"
	"gomall/app/bootstrap"
	. "gomall/app"
)

func main() {
	fmt.Printf("args:%v", os.Args)

	//解析启动参数
	var mode string
	flag.StringVar(&mode, "mode", "", "mode type")
	var level int
	flag.IntVar(&level, "level", 0, "debug level")
	flag.Parse()
	if mode == "" {
		flag.Usage()
		return
	}
	fmt.Printf("start mode:%s level:%d \n", mode, level)


	// log.SetFormatter(&log.TextFormatter{
	// 	FullTimestamp: true,
	// })
	// log.WithFields(log.Fields{
	// 	"animal": "walrus",
	// 	"number": 1,
	// 	"size":   10,
	// }).Info("A walrus appears")

	bootstrap.LoadConfig()

	bootstrap.SetupLog()
	GlobalLog.Info("Before app start")

	bootstrap.StartMysql()
	defer GlobalDb.Close()

	ginEngine := bootstrap.SetupRouters()

	fmt.Printf("app started %v \n", GlobalConfig)
	GlobalLog.Info("app started")

	ginEngine.Run()
}
