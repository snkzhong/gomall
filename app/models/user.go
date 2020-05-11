package models

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	gorm.Model
	UUID        uuid.UUID    `gorm:"column:uuid" json:"uuid"`
	Username    string       `gorm:"column:username" json:"username"`
	Password    string       `gorm:"column:password" json:"-"`
}