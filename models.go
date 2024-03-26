package main

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username       string
	HashedPassword string
}

type Consultant struct {
	gorm.Model
	Name                       string
	ExperienceInYears          int
	HourlySalaryRequirementSek int
	WorkProfiles               []WorkProfile `gorm:"foreignKey:ConsultantRefer"`
}

type Assignment struct {
	gorm.Model
	LocationID      uint
	Location        Location `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Shift           string
	HourlySalarySek int
	WorkProfiles    []WorkProfile `gorm:"many2many:workprofile_assignments;"`
}

type Location struct {
	gorm.Model
	Name         string
	Lon          float64
	Lat          float64
	Municipality string
	WorkProfiles []WorkProfile `gorm:"many2many:workprofile_locations;"`
}

type WorkProfile struct {
	gorm.Model
	ConsultantRefer uint
	Consultant      Consultant `gorm:"foreignKey:ConsultantRefer;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
