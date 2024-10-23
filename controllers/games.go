package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"table_top_api/models"
	"time"
)

type CreateGamesInput struct {
	Name     string    `json:"name" binding:"required"`
	Style    string    `json:"style" binding:"required"`
	Location string    `json:"location" binding:"required"`
	DateTime time.Time `json:"date_time" binding:"required"`
}

type UpdateGamesInput struct {
	Name     string    `json:"name"`
	Style    string    `json:"style"`
	Location string    `json:"location"`
	DateTime time.Time `json:"date_time"`
}

func FindGames(c *gin.Context) {
	var games []models.Game
	models.DB.Find(&games)

	c.JSON(http.StatusOK, gin.H{"data": games})
}

func gameReturn(id string) (*models.Game, error) {
	var game models.Game
	result := models.DB.Where("id = ?", id).First(&game)
	if result.Error != nil {
		return nil, result.Error
	}
	return &game, nil
}
func FindGameById(c *gin.Context) {
	game, err := gameReturn(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": *game})
}
func CreateGames(c *gin.Context) {
	var input CreateGamesInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// create game
	game := models.Game{
		Name:     input.Name,
		Style:    input.Style,
		Location: input.Location,
		DateTime: input.DateTime,
	}
	models.DB.Create(&game)

	c.JSON(http.StatusOK, gin.H{"data": game})
}

func UpdateGame(c *gin.Context) {
	game, err := gameReturn(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input UpdateGamesInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&game).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": game})
}

func DeleteGame(c *gin.Context) {
	game, err := gameReturn(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	models.DB.Delete(&game)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
