package domain

import "gorm.io/gorm"

type Transport struct {
	gorm.Model

	Domain    string `gorm:"type:varchar(255);not null;default:''" validate:"omitempty,max=255"`
	Transport string `gorm:"type:varchar(32);default:null" validate:"omitempty,max=32"`
}
