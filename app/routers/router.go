package routers

import (
	ProductRouters "gomall/app/modules/product/routers"

	"github.com/gin-gonic/gin"

	"gomall/app/global"
	"gomall/app/middlewares"
)

func StartRouters() {
	global.GlobalEngine = gin.Default()
	global.GlobalEngine.Use(middlewares.Test())

	ProductRouters.StartProductRouters(global.GlobalEngine)
}
