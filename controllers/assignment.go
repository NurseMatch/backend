package controllers

import (
	"backend/data"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterAssignmentEndpoints(e *gin.Engine) {
	e.GET("/assignment/:id", getAssignment)
	e.POST("/assignment", createAssignment)
	e.PUT("/assignment/:id", updateAssignment)
	e.DELETE("/assignment/:id", deleteAssignment)
}

func getAssignment(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var assignment data.Assignment
	id := c.Param("id")
	if err := db.Where("id = ?", id).First(&assignment).Error; err != nil {
		c.JSON(404, gin.H{
			"error": "No assignment found with given ID",
		})
		return
	}

	c.JSON(200, assignment)
}

func createAssignment(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var assignment data.Assignment
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

	c.JSON(201, assignment)
}

func updateAssignment(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var assignment data.Assignment
	id := c.Param("id")
	if err := db.Where("id = ?", id).First(&assignment).Error; err != nil {
		c.JSON(404, gin.H{
			"error": "No assignment found with given ID",
		})
		return
	}

	if err := c.ShouldBindJSON(&assignment); err != nil {
		c.JSON(400, gin.H{
			"error": "Bad Request",
		})
		return
	}

	result := db.Save(&assignment)

	if result.Error != nil {
		c.JSON(500, result.Error)
	}

	c.JSON(200, assignment)
}

func deleteAssignment(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var assignment data.Assignment
	id := c.Param("id")
	if err := db.Where("id = ?", id).Delete(&assignment).Error; err != nil {
		c.JSON(404, gin.H{
			"error": "No assignment found with given ID",
		})
		return
	}

	c.JSON(204, gin.H{})
}
