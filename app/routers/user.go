package routers

import (
	"github.com/gin-gonic/gin"
	"gomall/app/controllers"
)

func UserRouter(ginEngine *gin.Engine) {
	routerGroup := ginEngine.Group("")
	userRouter := routerGroup.Group("user")
	{
		userRouter.GET("register", controllers.UserRegister)
		userRouter.GET("login/:username", controllers.UserLogin)
		userRouter.GET("logout", controllers.UserLogout)
	}
}