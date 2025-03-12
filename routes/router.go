package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kelaasor-quiz/middleware"
)

func RouterHandler(server *gin.Engine) {
	server.POST("/login", LoginUserHandler)

	server.GET("/quizzes", GetAllQuizzesHandler)
	server.POST("/quizzes", AddNewQuizHandler)

	authRequire := server.Group("/")
	authRequire.Use(middleware.Authenticate)
	// questions
	authRequire.GET("/quizzes/:quizId/questions", GetQuestionsByQuizId)
	authRequire.POST("/quizzes/:quizId/questions", AddQuestionToQuizHandler)
	//answers
	authRequire.POST("/answers/:questionId", AddNewAnswerHandler)
	authRequire.GET("/answers/:questionId", GetAnswersByQuestionHandler)
	//reponse
	authRequire.POST("/submit/:questionId", SubmitAnswerHandler)

	authRequire.GET("/result", GetUserResultHandler)

}
