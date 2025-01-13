package routers

import (
	"github.com/gin-gonic/gin"
	"go-test/global"
	"go-test/user/model"
)

func ListAllUsers(c *gin.Context) {
	users := model.ListUsers(global.DB)
	c.JSON(200, gin.H{
		"data": users,
	})
}

func CreateUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "create",
	})
}
