package dao

import (
	"go-test/model"
	"gorm.io/gorm"
)

func CreateUser(user *model.User, db *gorm.DB) {
	db.Create(user)
}

func ListUsers(db *gorm.DB) []model.User {
	var users []model.User
	db.Model(&users).Find(&users)
	return users
}
