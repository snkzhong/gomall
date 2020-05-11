package bootstrap

import (
	"gomall/app/routers"
	"gomall/app/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRouters() *gin.Engine {
	ginEngine := gin.Default()
	ginEngine.Use(middlewares.Test())
	routers.Start(ginEngine)

	return ginEngine
}