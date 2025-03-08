package routes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kelaasor-quiz/models"
)

func SignupUserHandler(context *gin.Context) {
	user := models.User{}

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "email and passwor is required for signup"})
		return
	}

	tempUser, err := user.Signup()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"user": tempUser})

}

func LoginUserHandler(context *gin.Context) {
	user := models.User{}
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "phone and password is requierd for login"})
		return
	}

	accesstoken, err := models.ValidateUserCreadentials(user.Phone, user.Password)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"accessTtoken": accesstoken, "exp": time.Now().Add(time.Hour * 24).Unix()})

}
