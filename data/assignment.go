package data

import (
	"gorm.io/gorm"
)

type Assignment struct {
	gorm.Model
	LocationID      uint
	Location        Location `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Shift           string
	HourlySalarySek int
}
