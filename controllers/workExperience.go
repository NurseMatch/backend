package controllers

import (
	"backend/data"
	"backend/views"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func RegisterWorkExperienceEndpoints(e *gin.Engine) {
	e.POST("/workExperience", createWorkExperience)
	e.GET("/workExperience/:id", getWorkExperience)
	e.PUT("/workExperience/:id", updateWorkExperience)
	e.DELETE("/workExperience/:id", deleteWorkExperience)
}

func getWorkExperience(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	workExperienceID := c.Param("id")

	var workExperience data.WorkExperience
	if err := db.First(&workExperience, workExperienceID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "WorkExperience not found",
		})
		return
	}

	c.JSON(http.StatusOK, mapWorkExperienceView(workExperience))
}

func createWorkExperience(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var workExperienceView views.WorkExperience
	if err := c.ShouldBindJSON(&workExperienceView); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	newWorkExperience := mapWorkExperience(workExperienceView)

	if err := db.Create(&newWorkExperience).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create WorkExperience",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"WorkExperience_id": newWorkExperience.ID,
		"message":           "WorkExperience created successfully",
	})
}

func updateWorkExperience(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	workExperienceID := c.Param("id")

	var updatedWorkExperience views.WorkExperience
	if err := c.ShouldBindJSON(&updatedWorkExperience); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	var workExperience data.WorkExperience
	if err := db.First(&workExperience, workExperienceID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "WorkExperience not found",
		})
		return
	}

	workExperience.ConsultantID = updatedWorkExperience.ConsultantId
	workExperience.Title = updatedWorkExperience.Title
	workExperience.StartDate = updatedWorkExperience.StartDate
	workExperience.EndDate = updatedWorkExperience.EndDate
	workExperience.Description = updatedWorkExperience.Description

	if err := db.Save(&workExperience).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to update WorkExperience",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "WorkExperience updated successfully",
	})
}

func deleteWorkExperience(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	workExperienceID := c.Param("id")

	var workExperience data.WorkExperience
	if err := db.First(&workExperience, workExperienceID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "WorkExperience not found",
		})
		return
	}

	if err := db.Delete(&workExperience).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to delete WorkExperience",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "WorkExperience deleted successfully",
	})
}

func mapWorkExperienceView(experience data.WorkExperience) views.WorkExperience {
	return views.WorkExperience{
		Title:        experience.Title,
		Description:  experience.Description,
		StartDate:    experience.StartDate,
		EndDate:      experience.EndDate,
		ConsultantId: experience.ConsultantID,
	}
}

func mapWorkExperience(experience views.WorkExperience) data.WorkExperience {
	return data.WorkExperience{
		Title:        experience.Title,
		Description:  experience.Description,
		StartDate:    experience.StartDate,
		EndDate:      experience.EndDate,
		ConsultantID: experience.ConsultantId,
	}
}
