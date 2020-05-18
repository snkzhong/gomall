package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"gomall/app/bootstrap"
	"gomall/app/global"
	"gomall/app/routers"
	"gomall/scripts"
)

func main() {
	fmt.Printf("args:%v \n", os.Args)

	//解析启动参数
	/*
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
	*/

	bootstrap.LoadConfig()

	bootstrap.SetupLog()
	global.GlobalLog.Info("Before app start")

	bootstrap.StartMysql()
	defer global.GlobalDb.Close()

	// bootstrap.SetupRouters()
	routers.StartRouters()

	var genModel string
	flag.StringVar(&genModel, "genModel", "", "Generate Model")
	flag.Parse()
	if genModel != "" {
		modelTable := strings.Split(genModel, ":")
		if len(modelTable) == 2 {
			scripts.GenModel(modelTable[0], modelTable[1])
		}
		return
	}

	// fmt.Printf("categorys: %v \n", services.GetAllCategory())
	// services.GetAllCategory()
	global.GlobalLog.Info("app started")
	bootstrap.Startup()
}
