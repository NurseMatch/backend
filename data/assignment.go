package data

import "gorm.io/gorm"

type Assignment struct {
	gorm.Model
	DisplayName          string
	Description          string
	HourlyRate           int
	EducationRequirement string
	Role                 string
	Location             string
	Worker               Worker `gorm:"foreignKey:ID"`
}
