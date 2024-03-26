package data

import (
	"gorm.io/gorm"
)

type Location struct {
	gorm.Model
	Name         string
	Lon          float64
	Lat          float64
	Municipality string
	WorkProfiles []WorkProfile `gorm:"many2many:workprofile_locations;"`
}
