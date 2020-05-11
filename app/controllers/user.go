package controllers

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"gomall/app/services"
	"gomall/app/models"
	// . "gomall/app"
)

func UserRegister(context *gin.Context) {
	username := context.Query("username")
	user := &models.User{Username:username, Password:"123456"}
	r,err := services.Register(*user)
	fmt.Printf("username:%s \n", username)
	// r, err := GlobalDb.DB().Exec("INSERT INTO user(username, password) VALUES(?,?)","ZhongQin","123456")
	if err != nil {
		fmt.Println("error:%v", err)
		context.JSON(http.StatusOK, gin.H{
			"op": "Register",
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"op": "Register",
			"result": r,
		})
	}
}

func UserLogin(context *gin.Context) {
	username := context.Param("username")
	user, _ := services.FindByUsername(username)
	context.JSON(http.StatusOK, gin.H{
		"op": "Login",
		"user": user,
	})
}

func UserLogout(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"user": "Logout",
	})
}