package routers

import (
	"github.com/gin-gonic/gin"
	"go-test/dao"
	"go-test/global"
	"go-test/model"
)

func ListAllUsers(c *gin.Context) {
	users := dao.ListUsers(global.DB)
	c.JSON(200, gin.H{
		"data": users,
	})
}

func CreateUser(c *gin.Context) {
	type UserCreateDTO struct {
		Name     string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
		Age      int    `json:"age"`
		Email    string `json:"email"`
	}
	var req UserCreateDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	m := model.User{
		Name:     req.Name,
		Password: req.Password,
		Age:      req.Age,
		Email:    req.Email,
		Status:   1,
	}
	dao.CreateUser(&m, global.DB)
	c.JSON(200, req)
}
