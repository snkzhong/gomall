package controllers

import (
	"gomall/app/modules/product/services"
	"net/http"

	"gomall/app/modules/product/forms"

	"fmt"

	"gomall/app/libs"

	"github.com/gin-gonic/gin"
)

func CategoryList(context *gin.Context) {
	var categoryForm forms.CategoryForm
	err := context.ShouldBind(&categoryForm)
	fmt.Printf("\n \t categoryForm: %v \n", categoryForm)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// context.JSON(http.StatusOK, gin.H{
	// 	"categorys": services.GetAllCategory(),
	// })

	libs.NewResponse().OkWithData(gin.H{
		"categorys": services.GetAllCategory(),
	}, context)
}
