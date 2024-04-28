package data

import (
	"gorm.io/gorm"
)

type Assignment struct {
	gorm.Model
	Shift           string
	HourlySalarySek int
	AssignmentRefer uint
}
