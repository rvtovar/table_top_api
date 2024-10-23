package routes

import (
	"github.com/gin-gonic/gin"
	"table_top_api/controllers"
)

func RegisterGameRoutes(c *gin.Engine) {

	c.GET("/games", controllers.FindGames)
	c.POST("/games", controllers.CreateGames)
	c.GET("games/:id", controllers.FindGameById)
	c.PATCH("games/:id", controllers.UpdateGame)
	c.DELETE("games/:id", controllers.DeleteGame)

}
