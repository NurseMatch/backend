package controllers

import (
	"backend/data"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterWorkplaceEndpoints(e *gin.Engine) {
	e.POST("/workplace", createWorkplace)

}

func createWorkplace(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var assignment data.Workplace
	if err := c.ShouldBindJSON(&assignment); err != nil {
		c.JSON(400, gin.H{
			"error": "Bad Request",
		})
		return
	}

	result := db.Create(&assignment)

	if result.Error != nil {
		c.JSON(500, result.Error)
	}

	c.JSON(201, assignment.ID)
}
