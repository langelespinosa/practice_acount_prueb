package domain

import "gorm.io/gorm"

type Aliases struct {
	gorm.Model

	Local  string `gorm:"type:varchar(255);not null;default:'';" validate:"required,email,max=255"`
	Remoto string `gorm:"type:text;default:null;" validate:"omitempty,max=65535"`
}
