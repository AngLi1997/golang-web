package model

import (
	"gorm.io/gorm"
)

type User struct {
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      int    `json:"age"`
	Job      string `json:"job"`
	Status   int    `json:"status"`
	gorm.Model
}

func (User) TableName() string {
	return "user"
}
