package data

import (
	"gorm.io/gorm"
)

type Consultant struct {
	gorm.Model
	Name                       string
	ExperienceInYears          int
	HourlySalaryRequirementSek int
	WorkProfiles               []WorkProfile `gorm:"foreignKey:ConsultantRefer"`
}
