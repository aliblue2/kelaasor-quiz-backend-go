package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kelaasor-quiz/models"
)

func GetQuestionsByQuizId(context *gin.Context) {
	quizId, err := strconv.ParseInt(context.Param("quizId"), 10, 64)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	_, err = models.GetQuizWithId(quizId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	questions, err := models.GetQuestionByQuizId(quizId)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"questions": questions})

}

func AddQuestionToQuizHandler(context *gin.Context) {
	quizId, err := strconv.ParseInt(context.Param("quizId"), 10, 64)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	_, err = models.GetQuizWithId(quizId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	question := models.Question{}
	context.ShouldBindJSON(&question)
	question.QuizId = quizId

	questionId, err := question.AddQuestionToQuiz()

	if err != nil {
		context.JSON(http.StatusConflict, gin.H{"message": err.Error()})
		return
	}

	question.Id = questionId
	context.JSON(http.StatusCreated, gin.H{"message": "successfully created.!", "question": question})

}
