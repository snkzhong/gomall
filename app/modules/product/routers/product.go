package routers

import (
	"gomall/app/modules/product/controllers"

	"github.com/gin-gonic/gin"
)

func StartProductRouters(ginEngine *gin.Engine) {
	routerGroup := ginEngine.Group("")
	categoryRouter := routerGroup.Group("product/category")
	{
		categoryRouter.GET("list", controllers.CategoryList)
		categoryRouter.POST("add", controllers.CategoryList)
	}
}
