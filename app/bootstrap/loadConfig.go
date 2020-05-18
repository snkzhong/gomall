package bootstrap

import (
	"fmt"

	"gomall/app/global"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func LoadConfig() {
	v := viper.New()
	v.SetConfigFile(global.DefaultConfigFile)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		if err := v.Unmarshal(&global.GlobalConfig); err != nil {
			fmt.Println(err)
		}
		fmt.Printf("mysql:%v, redis:%v \n", global.GlobalConfig.Mysql, global.GlobalConfig.Redis)
	})

	if err := v.Unmarshal(&global.GlobalConfig); err != nil {
		fmt.Println(err)
	}
}
