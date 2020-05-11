package app

import (
	"gomall/app/config"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"os"
)

const DefaultConfigFile = "config/config.yaml"

var (
	GlobalConfig config.Config
	GlobalDb *gorm.DB
	GlobalLog *logrus.Logger
	GlobalSqlLogFile *os.File
)
