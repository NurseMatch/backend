package controllers

import (
	"backend/data"
	"backend/views"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func RegisterWorkerEndpoints(e *gin.Engine) {
	e.POST("/worker", createWorker)
	e.GET("/worker/:id", getWorker)
	e.PUT("/worker/:id", updateWorker)
	e.DELETE("/worker/:id", deleteWorker)
}

func getWorker(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	workerID := c.Param("id")

	var worker data.Worker
	if err := db.Joins("left join work_experiences on work_experiences.id = workers.id").First(&worker, workerID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Worker not found",
		})
		return
	}

	workerView := views.Worker{
		Name:           worker.Name,
		WorkExperience: mapWorkExperienceView(worker.WorkExperience),
		Education:      worker.Education,
		Role:           worker.Role,
		Email:          worker.Email,
		Phone:          worker.Phone,
		ProfileImage:   worker.ProfileImage,
		JobPreferences: worker.JobPreferences,
		HourlyRate:     worker.HourlyRate,
		Description:    worker.Description,
		Location:       worker.Location,
		IDImage:        worker.IDImage,
	}

	c.JSON(http.StatusOK, workerView)
}

func createWorker(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var workerView views.Worker
	if err := c.ShouldBindJSON(&workerView); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	var existingWorker data.Worker
	if err := db.Where("email = ?", workerView.Email).First(&existingWorker).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": "Email already exists",
		})
		return
	}

	newWorker := data.Worker{
		Name:           workerView.Name,
		WorkExperience: mapWorkExperience(workerView.WorkExperience),
		Education:      workerView.Education,
		Role:           workerView.Role,
		Email:          workerView.Email,
		Phone:          workerView.Phone,
		ProfileImage:   workerView.ProfileImage,
		JobPreferences: workerView.JobPreferences,
		HourlyRate:     workerView.HourlyRate,
		Description:    workerView.Description,
		Location:       workerView.Location,
		IDImage:        workerView.IDImage,
	}

	if err := db.Create(&newWorker).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create worker",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"worker_id": newWorker.ID,
		"message":   "Worker created successfully",
	})
}

func updateWorker(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	workerId := c.Param("id")

	var updatedWorker views.Worker
	if err := c.ShouldBindJSON(&updatedWorker); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
		})
		return
	}

	var worker data.Worker
	if err := db.First(&worker, workerId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Worker not found",
		})
		return
	}

	worker.Name = updatedWorker.Name
	worker.WorkExperience = mapWorkExperience(updatedWorker.WorkExperience)
	worker.Education = updatedWorker.Education
	worker.Role = updatedWorker.Role
	worker.Email = updatedWorker.Email
	worker.Phone = updatedWorker.Phone
	worker.ProfileImage = updatedWorker.ProfileImage
	worker.JobPreferences = updatedWorker.JobPreferences
	worker.HourlyRate = updatedWorker.HourlyRate
	worker.Description = updatedWorker.Description
	worker.Location = updatedWorker.Location
	worker.IDImage = updatedWorker.IDImage

	if err := db.Save(&worker).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update worker",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Worker updated successfully",
	})
}

func deleteWorker(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	workerId := c.Param("id")

	var worker data.Worker
	if err := db.First(&worker, workerId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Worker not found",
		})
		return
	}

	if err := db.Delete(&worker).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete worker",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Worker deleted successfully",
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
