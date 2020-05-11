package bootstrap

import (
	"fmt"
	"os"
	// "log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	. "gomall/app"
	"gomall/app/libs"
)

func StartMysql() {
	mysqlConfig := GlobalConfig.Mysql
	var url string = fmt.Sprintf("%s:%s@%s/%s?%s", mysqlConfig.Username, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Database, mysqlConfig.Config)
	// fmt.Printf("mysql url: %s", url)
	db,err := gorm.Open("mysql", url)
	
	if err != nil {
		fmt.Printf("mysql error: %s", err)
		os.Exit(0)
	} else {
		GlobalDb = db
		GlobalDb.DB().SetMaxIdleConns(mysqlConfig.MaxIdleConns)
		GlobalDb.DB().SetMaxOpenConns(mysqlConfig.MaxOpenConns)
		GlobalDb.SingularTable(mysqlConfig.SingularTable)
		GlobalDb.LogMode(mysqlConfig.LogMode)

		if GlobalConfig.Log.Sql.Debug {
			logger := &libs.Logger{}
			GlobalDb.SetLogger(logger)
			f, _ := os.OpenFile(GlobalConfig.Log.Sql.File, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
			GlobalSqlLogFile = f
		}
	}
}