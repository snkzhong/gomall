package middlewares

import (
	"github.com/gin-gonic/gin"
	// "fmt"
)

func Test() gin.HandlerFunc {
	return func(context *gin.Context) {
		// request := context.Request
		// fmt.Printf("middlewares test:%v \n", request)

		context.Next()
	}
}