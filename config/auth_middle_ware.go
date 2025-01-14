package config

import (
	"github.com/gin-gonic/gin"
	"go-test/common"
	"net/http"
)

func TokenAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.GetHeader("Authorization")
		if token == "" {
			context.JSON(http.StatusUnauthorized, common.Resp{Msg: "未认证"})
			context.Abort()
		}
		// 校验token
		context.Next()
	}
}
