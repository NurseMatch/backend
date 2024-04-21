package data

import (
	"gorm.io/gorm"
)

type Workplace struct {
	gorm.Model
	Name         string `gorm:"type:varchar(100);uniqueIndex"`
	Municipality string
	Assignments  []Assignment `gorm:"foreignKey:AssignmentRefer;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
