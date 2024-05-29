package data

import (
	"gorm.io/gorm"
	"time"
)

type Worker struct {
	gorm.Model
	Name           string
	WorkExperience WorkExperience `gorm:"foreignKey:ID"`
	Education      string
	Role           string
	Email          string
	Phone          string
	ProfileImage   string
	JobPreferences string
	HourlyRate     float64
	Description    string
	Location       string
	IDImage        string
}

type WorkExperience struct {
	gorm.Model
	Title       string
	Description string
	StartDate   time.Time
	EndDate     time.Time
}
