package controllers

import (
	"backend/data"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterAccountEndpoints(e *gin.Engine) {
	e.POST("/createUser", createUser)
}

func createUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var user data.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{
			"error": "Bad Request",
		})
		return
	}

	result := db.Create(&user)

	if result.Error != nil {
		c.JSON(500, result.Error)
	}

	c.JSON(201, user.ID)
}
