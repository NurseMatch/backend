package data

import (
	"gorm.io/gorm"
	"time"
)

type Consultant struct {
	gorm.Model
	Name            string
	Education       string
	Role            string
	Email           string `gorm:"uniqueIndex"`
	Phone           string
	ProfileImage    string
	JobPreferences  string
	HourlyRate      float64
	Description     string
	Location        string
	IDImage         string
	WorkExperiences []WorkExperience `gorm:"foreignKey:ConsultantID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type WorkExperience struct {
	gorm.Model
	Title        string
	Description  string
	StartDate    time.Time
	EndDate      time.Time
	ConsultantID uint `gorm:"foreignKey:ConsultantID;references:ID"`
}

func (w *WorkExperience) BeforeCreate(tx *gorm.DB) (err error) {
	// Ensure ConsultantID is valid before creating a WorkExperience record
	var consultant Consultant
	if err = tx.First(&consultant, w.ConsultantID).Error; err != nil {
		return err
	}
	return nil
}
