package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Start(ginEngine *gin.Engine) {
	fmt.Println("routers Start")

	IndexRouter(ginEngine)
	UserRouter(ginEngine)
}