package config

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func TokenAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.GetHeader("Authorization")
		if token == "" {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "token is required"})
		}
		context.Next()
	}
}
