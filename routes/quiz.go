package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kelaasor-quiz/models"
)

func GetAllQuizzesHandler(context *gin.Context) {
	quizzes, err := models.GetAllQuizzes()

	if err != nil {
		context.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"quizzes": quizzes})
}

func AddNewQuizHandler(context *gin.Context) {
	quiz := models.Quiz{}

	err := context.ShouldBindJSON(&quiz)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	_, err = quiz.AddNewQuiz()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "quiz successfully created"})

}
