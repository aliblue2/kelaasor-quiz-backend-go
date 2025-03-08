package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kelaasor-quiz/models"
)

func AddNewAnswerHandler(context *gin.Context) {
	questionId, err := strconv.ParseInt(context.Param("questionId"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	_, err = models.GetQuestionById(questionId)

	if err != nil {
		context.JSON(http.StatusConflict, gin.H{"message": err.Error()})
		return
	}

	answer := models.Answer{}

	err = context.ShouldBindJSON(&answer)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	answer.QuestionId = questionId
	_, err = answer.AddNewAnswerToQuestions()

	if err != nil {
		context.JSON(http.StatusConflict, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "answer successfully added.!"})

}

func GetAnswersByQuestionHandler(context *gin.Context) {
	questionId, err := strconv.ParseInt(context.Param("questionId"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	_, err = models.GetQuestionById(questionId)

	if err != nil {
		context.JSON(http.StatusConflict, gin.H{"message": err.Error()})
		return
	}

	answers, err := models.GetAllAnswersByQuestionId(questionId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"answers": answers})

}
