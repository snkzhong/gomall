package global

import (
	"gomall/app/config"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

const DefaultConfigFile = "config/config.yaml"

var (
	GlobalEngine     *gin.Engine
	GlobalConfig     config.Config
	GlobalDb         *gorm.DB
	GlobalLog        *logrus.Logger
	GlobalSqlLogFile *os.File
)
