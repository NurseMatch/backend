package controllers

import (
	"backend/data"
	"backend/views"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func RegisterConsultantEndpoints(e *gin.Engine) {
	e.POST("/consultant", createConsultant)
	e.GET("/consultant/:id", getConsultant)
	e.PUT("/consultant/:id", updateConsultant)
	e.DELETE("/consultant/:id", deleteConsultant)
}

func getConsultant(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	consultantID := c.Param("id")

	var consultant data.Consultant
	if err := db.First(&consultant, consultantID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Consultant not found",
		})
		return
	}

	c.JSON(http.StatusOK, consultant)
}

func createConsultant(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var consultantView views.Consultant
	if err := c.ShouldBindJSON(&consultantView); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
		})
		return
	}

	var existingUser data.User
	if err := db.Where("email = ?", consultantView.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": "Email already exists",
		})
		return
	}

	newConsultant := data.Consultant{
		Name:              consultantView.Name,
		ExperienceInYears: consultantView.Experience,
		Education:         consultantView.Education,
		Role:              consultantView.Role,
		Email:             consultantView.Email,
		Phone:             consultantView.Phone,
		ProfileImage:      consultantView.ProfileImage,
		JobPreferences:    consultantView.JobPreferences,
		HourlyRate:        consultantView.HourlyRate,
		Description:       consultantView.Description,
		Location:          consultantView.Location,
		IDImage:           consultantView.IDImage,
	}

	if err := db.Create(&newConsultant).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"consultant_id": newConsultant.ID,
		"message":       "Consultant created successfully",
	})
}

func updateConsultant(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	consultantID := c.Param("id")

	var consultant data.Consultant
	if err := db.First(&consultant, consultantID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Consultant not found",
		})
		return
	}

	var updatedConsultant views.Consultant
	if err := c.ShouldBindJSON(&updatedConsultant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
		})
		return
	}

	consultant.Name = updatedConsultant.Name
	consultant.ExperienceInYears = updatedConsultant.Experience
	consultant.Education = updatedConsultant.Education
	consultant.Role = updatedConsultant.Role
	consultant.Email = updatedConsultant.Email
	consultant.Phone = updatedConsultant.Phone
	consultant.ProfileImage = updatedConsultant.ProfileImage
	consultant.JobPreferences = updatedConsultant.JobPreferences
	consultant.HourlyRate = updatedConsultant.HourlyRate
	consultant.Description = updatedConsultant.Description
	consultant.Location = updatedConsultant.Location
	consultant.IDImage = updatedConsultant.IDImage

	if err := db.Save(&consultant).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update consultant",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Consultant updated successfully",
	})
}

func deleteConsultant(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	consultantID := c.Param("id")

	var consultant data.Consultant
	if err := db.First(&consultant, consultantID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Consultant not found",
		})
		return
	}

	if err := db.Delete(&consultant).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete consultant",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Consultant deleted successfully",
	})
}
