package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Product(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "hello product!",
	})
}
