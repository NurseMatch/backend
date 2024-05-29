package views

import (
	"github.com/google/uuid"
	"time"
)

type Worker struct {
	Name           string         `json:"name"`
	WorkExperience WorkExperience `json:"workExperience"`
	Education      string         `json:"education"`
	Role           string         `json:"role"`
	Email          string         `json:"email"`
	Phone          string         `json:"phone"`
	ProfileImage   string         `json:"profile_image"`
	JobPreferences string         `json:"job_preferences"`
	HourlyRate     float64        `json:"hourly_rate"`
	Description    string         `json:"description"`
	Location       string         `json:"location"`
	IDImage        string         `json:"id_image"`
}

type WorkExperience struct {
	Id          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"startDate"`
	EndDate     time.Time `json:"endDate"`
}
