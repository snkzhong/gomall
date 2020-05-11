package routers

import (
	"github.com/gin-gonic/gin"
	"gomall/app/controllers"
)

func IndexRouter(ginEngine *gin.Engine) {
	ginEngine.GET("/", controllers.Index)
}