package routes

import (
	"github.com/gin-gonic/gin"
	"table_top_api/controllers"
	"table_top_api/middleware"
)

func RegisterGameRoutes(c *gin.Engine) {

	c.GET("/games", controllers.FindGames)
	c.GET("games/:id", controllers.FindGameById)

	authenticated := c.Group("/")
	authenticated.Use(middleware.Authenticate)
	authenticated.POST("/games", controllers.CreateGames)
	authenticated.PATCH("games/:id", controllers.UpdateGame)
	authenticated.DELETE("games/:id", controllers.DeleteGame)

	authenticated.POST("/games/:id/register", controllers.RegisterForGame)
	authenticated.DELETE("/games/:id/register", controllers.CancelRegistration)

}
