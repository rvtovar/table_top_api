package routes

import (
	"github.com/gin-gonic/gin"
	"table_top_api/controllers"
)

func RegisterUserRoutes(c *gin.Engine) {

	c.POST("/signup", controllers.CreateUser)
	c.POST("/login", controllers.Login)

}
