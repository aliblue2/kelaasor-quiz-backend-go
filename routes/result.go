package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kelaasor-quiz/models"
)

func GetUserResultHandler(context *gin.Context) {
	userId := context.GetInt64("userId")

	personalTypes, err := models.GetUserResult(userId)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": personalTypes})

}
