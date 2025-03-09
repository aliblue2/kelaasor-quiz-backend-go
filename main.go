package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kelaasor-quiz/db"
	"github.com/kelaasor-quiz/routes"
)

func main() {

	db.DatabaseConnection()
	server := gin.Default()

	routes.RouterHandler(server)

	server.Run(":8000")

}
