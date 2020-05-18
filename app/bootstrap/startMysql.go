package bootstrap

import (
	"fmt"
	"os"

	// "log"
	"gomall/app/global"
	"gomall/app/libs"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func StartMysql() {
	mysqlConfig := global.GlobalConfig.Mysql
	var url string = fmt.Sprintf("%s:%s@%s/%s?%s", mysqlConfig.Username, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Database, mysqlConfig.Config)
	// fmt.Printf("mysql url: %s", url)
	db, err := gorm.Open("mysql", url)

	if err != nil {
		fmt.Printf("mysql error: %s", err)
		os.Exit(0)
	} else {
		global.GlobalDb = db
		global.GlobalDb.DB().SetMaxIdleConns(mysqlConfig.MaxIdleConns)
		global.GlobalDb.DB().SetMaxOpenConns(mysqlConfig.MaxOpenConns)
		global.GlobalDb.SingularTable(mysqlConfig.SingularTable)
		global.GlobalDb.LogMode(mysqlConfig.LogMode)

		if global.GlobalConfig.Log.Sql.Debug {
			logger := &libs.Logger{}
			global.GlobalDb.SetLogger(logger)
			f, _ := os.OpenFile(global.GlobalConfig.Log.Sql.File, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
			global.GlobalSqlLogFile = f
		}
	}
}
