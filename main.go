package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"table_top_api/models"
	"table_top_api/routes"
)

func main() {
	server := gin.Default()
	server.GET("/", welcome)
	routes.RegisterGameRoutes(server)
	routes.RegisterUserRoutes(server)

	models.InitDB()
	server.Run(":8080")
}

func welcome(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Welcome to Game Night!"})
}
