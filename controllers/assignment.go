package controllers

import (
	"backend/data"
	"backend/views"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func RegisterAssignmentEndpoints(e *gin.Engine) {
	e.POST("/assignment", createAssignment)
	e.GET("/assignment/:id", getAssignment)
	e.PUT("/assignment/:id", updateAssignment)
	e.DELETE("/assignment/:id", deleteAssignment)
}

func getAssignment(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	assignmentId := c.Param("id")

	var assignment data.Assignment
	if err := db.First(&assignment, assignmentId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Assignment not found",
		})
		return
	}

	assignmentView := views.Assignment{
		DisplayName:          assignment.DisplayName,
		Description:          assignment.Description,
		HourlyRate:           assignment.HourlyRate,
		EducationRequirement: assignment.EducationRequirement,
		Role:                 assignment.Role,
		Location:             assignment.Location,
	}

	c.JSON(http.StatusOK, assignmentView)
}

func createAssignment(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var assignmentView views.Assignment
	if err := c.ShouldBindJSON(&assignmentView); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	newAssignment := data.Assignment{
		DisplayName:          assignmentView.DisplayName,
		Description:          assignmentView.Description,
		HourlyRate:           assignmentView.HourlyRate,
		EducationRequirement: assignmentView.EducationRequirement,
		Role:                 assignmentView.Role,
		Location:             assignmentView.Location,
	}

	if err := db.Create(&newAssignment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create assignment",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"assignment_id": newAssignment.ID,
		"message":       "Assignment created successfully",
	})
}

func updateAssignment(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	assignmentId := c.Param("id")

	var updatedAssignment views.Assignment
	if err := c.ShouldBindJSON(&updatedAssignment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
		})
		return
	}

	var assignment data.Assignment
	if err := db.First(&assignment, assignmentId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Assignment not found",
		})
		return
	}

	assignment.Description = updatedAssignment.Description
	assignment.DisplayName = updatedAssignment.DisplayName
	assignment.Role = updatedAssignment.Role
	assignment.HourlyRate = updatedAssignment.HourlyRate
	assignment.Location = updatedAssignment.Location
	assignment.EducationRequirement = updatedAssignment.EducationRequirement

	if err := db.Save(&assignment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update assignment",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Assignment updated successfully",
	})
}

func deleteAssignment(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	assignmentId := c.Param("id")

	var assignment data.Assignment
	if err := db.First(&assignment, assignmentId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Assignment not found",
		})
		return
	}

	if err := db.Delete(&assignment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete assignment",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Assignment deleted successfully",
	})
}
