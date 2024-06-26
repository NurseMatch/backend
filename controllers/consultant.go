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
	if err := db.Joins("left join work_experiences on work_experiences.id = consultants.id").First(&consultant, consultantID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Consultant not found",
		})
		return
	}

	consultantView := views.Consultant{
		Name:           consultant.Name,
		WorkExperience: mapWorkExperienceView(consultant.WorkExperience),
		Education:      consultant.Education,
		Role:           consultant.Role,
		Email:          consultant.Email,
		Phone:          consultant.Phone,
		ProfileImage:   consultant.ProfileImage,
		JobPreferences: consultant.JobPreferences,
		HourlyRate:     consultant.HourlyRate,
		Description:    consultant.Description,
		Location:       consultant.Location,
		IDImage:        consultant.IDImage,
	}

	c.JSON(http.StatusOK, consultantView)
}

func createConsultant(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var consultantView views.Consultant
	if err := c.ShouldBindJSON(&consultantView); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	var existingConsultant data.Consultant
	if err := db.Where("email = ?", consultantView.Email).First(&existingConsultant).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": "Email already exists",
		})
		return
	}

	newConsultant := data.Consultant{
		Name:           consultantView.Name,
		WorkExperience: mapWorkExperience(consultantView.WorkExperience),
		Education:      consultantView.Education,
		Role:           consultantView.Role,
		Email:          consultantView.Email,
		Phone:          consultantView.Phone,
		ProfileImage:   consultantView.ProfileImage,
		JobPreferences: consultantView.JobPreferences,
		HourlyRate:     consultantView.HourlyRate,
		Description:    consultantView.Description,
		Location:       consultantView.Location,
		IDImage:        consultantView.IDImage,
	}

	if err := db.Create(&newConsultant).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create consultant",
			"message": err.Error(),
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

	var updatedConsultant views.Consultant
	if err := c.ShouldBindJSON(&updatedConsultant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
		})
		return
	}

	var consultant data.Consultant
	if err := db.First(&consultant, consultantID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Consultant not found",
		})
		return
	}

	consultant.Name = updatedConsultant.Name
	consultant.WorkExperience = mapWorkExperience(updatedConsultant.WorkExperience)
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

func mapWorkExperienceView(experience data.WorkExperience) views.WorkExperience {
	return views.WorkExperience{
		Title:       experience.Title,
		Description: experience.Description,
		StartDate:   experience.StartDate,
		EndDate:     experience.EndDate,
	}
}

func mapWorkExperience(experience views.WorkExperience) data.WorkExperience {
	return data.WorkExperience{
		Title:       experience.Title,
		Description: experience.Description,
		StartDate:   experience.StartDate,
		EndDate:     experience.EndDate,
	}
}
