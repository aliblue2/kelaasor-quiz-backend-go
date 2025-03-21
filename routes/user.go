package routes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kelaasor-quiz/models"
	"github.com/kelaasor-quiz/utils"
)

const temprorayPassword = "kelaasor.com/kelaasor-quiz"

func LoginUserHandler(context *gin.Context) {
	user := models.User{}
	err := context.ShouldBindJSON(&user)
	user.Password = temprorayPassword

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "phone and password is requierd for login"})
		return
	}

	userId, err := user.Signup()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	user.Id = userId

	accesstoken, err := utils.GenerateToken(user.Phone, user.Password, userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"accessTtoken": accesstoken, "exp": time.Now().Add(time.Hour * 24).Unix()})

}
