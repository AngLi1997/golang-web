package config

import (
	"github.com/gin-gonic/gin"
	"go-test/routers"
)

func InitRouters() {
	r := gin.Default()

	userGroup := r.Group("/user")
	userGroup.GET("/list_users", routers.ListAllUsers)
	userGroup.POST("/create_user", routers.CreateUser)

	err := r.Run("127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
}
