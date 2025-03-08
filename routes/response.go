package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kelaasor-quiz/models"
)

func SubmitAnswerHandler(context *gin.Context) {
	userId := context.GetInt64("userId")

	response := models.Response{}

	err := context.ShouldBindJSON(&response)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	response.UserId = userId
	id, err := response.Submit()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	response.Id = id

	context.JSON(http.StatusOK, gin.H{"message": "reponse submitted.!"})

}
