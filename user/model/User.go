package model

import "gorm.io/gorm"

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      int    `json:"age"`
	Status   int    `json:"status"`
	gorm.Model
}

func CreateInitUsers(db *gorm.DB) {
	err := db.AutoMigrate(&User{})
	if err != nil {
		return
	}
	users := []User{
		{Name: "user1", Email: "user1@qq.com", Password: "123456", Age: 18, Status: 1},
		{Name: "user2", Email: "user2@qq.com", Password: "123456", Age: 18, Status: 1},
		{Name: "user3", Email: "user3@qq.com", Password: "123456", Age: 18, Status: 1},
	}
	db.Create(&users)
}

func ListUsers(db *gorm.DB) []User {
	var users []User
	db.Model(&users).Find(&users)
	return users
}
