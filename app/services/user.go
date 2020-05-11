package services

import (
	"gomall/app/models"
	"gomall/app/utils"
	. "gomall/app"
	uuid "github.com/satori/go.uuid"
)

func Register(user models.User) (result models.User, err error) {
	user.Password = utils.Md5v([]byte(user.Password))
	user.UUID = uuid.NewV4()
	err = GlobalDb.Create(&user).Error

	return user, err
}

func FindByUsername(username string) (user models.User, err error) {
	GlobalDb.Where("username=?", username).First(&user)

	return user, err
}