package controllers

import (
	"github.com/gin-gonic/gin"
	"table_top_api/models"
	"table_top_api/utils"
)

type CreateUserInput struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func CreateUser(c *gin.Context) {
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	//create user
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		c.JSON(400, gin.H{"error": "Error hashing password"})
		return
	}

	user := models.User{
		Username: input.Username,
		Email:    input.Email,
		Password: hashedPassword,
	}

	if err := models.DB.Create(&user).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": "User created"})
}

func Login(c *gin.Context) {
	var input models.User
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = input.ValidateCreds()
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	token, err := utils.GenerateToken(input.Email, input.ID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"token": token})
}
