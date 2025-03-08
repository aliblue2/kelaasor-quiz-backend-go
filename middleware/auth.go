package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kelaasor-quiz/utils"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("token")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "missed token"})
		return
	}

	userId, err := utils.ValidateToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	context.Set("userId", userId)
}
