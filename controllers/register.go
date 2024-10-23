package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterForGame(c *gin.Context) {
	userId := c.GetInt64("uid")

	game, err := gameReturn(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not find game"})
		return
	}

	err = game.Register(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not register for game"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Successfully registered for game"})
}

func CancelRegistration(c *gin.Context) {
	userId := c.GetInt64("uid")

	game, err := gameReturn(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not find game"})
		return
	}

	err = game.CancelRegistration(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not cancel registration"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully cancelled registration"})
}
