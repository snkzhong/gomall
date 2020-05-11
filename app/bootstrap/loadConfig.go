package bootstrap

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/fsnotify/fsnotify"
	. "gomall/app"
)

func LoadConfig() {
	v := viper.New()
	v.SetConfigFile(DefaultConfigFile)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		if err := v.Unmarshal(&GlobalConfig); err != nil {
			fmt.Println(err)
		}
		fmt.Printf("mysql:%v, redis:%v \n", GlobalConfig.Mysql, GlobalConfig.Redis)
	})

	if err := v.Unmarshal(&GlobalConfig); err != nil {
		fmt.Println(err)
	}
}