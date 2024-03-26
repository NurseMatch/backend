package data

import (
	"gorm.io/gorm"
)

type WorkProfile struct {
	gorm.Model
	ConsultantRefer uint
	Consultant      Consultant `gorm:"foreignKey:ConsultantRefer;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
