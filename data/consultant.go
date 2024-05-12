package data

import "gorm.io/gorm"

type Consultant struct {
	gorm.Model
	Name              string
	ExperienceInYears int
	Education         string
	Role              string
	Email             string
	Phone             string
	ProfileImage      string
	JobPreferences    string
	HourlyRate        float64
	Description       string
	Location          string
	IDImage           string
}
