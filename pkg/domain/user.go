package domain

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model

	Userid               int       `gorm:"type:int;not null;default:65000"`
	Login                string    `gorm:"type:varchar(255);not null;default:''" validate:"min=4,max=255"`
	Email                string    `gorm:"type:varchar(255);not null;default:'';unique" validate:"omitempty,email"`
	Password             string    `gorm:"type:varchar(60);default:null" validate:"omitempty,min=8,containsany=!¡@#$%&=¿?*-_."`
	Maildir              string    `gorm:"type:mediumtext;not null"`
	Identificacion       string    `gorm:"type:varchar(255);not null;default:''" validate:"omitempty,max=255"`
	Grupo                string    `gorm:"type:varchar(255);not null;default:''" validate:"omitempty,max=255"`
	Dominio              int       `gorm:"type:int;not null;default:1" validate:"omitempty,gte=1"`
	Quota                int32     `gorm:"type:bigint;not null;default:30000000" validate:"omitempty,gte=0"`
	Policyd_messagequota uint      `gorm:"type:int;not null;default:100" validate:"omitempty,gte=0"`
	Policyd_messagetally int       `gorm:"type:int;not null;default:0" validate:"omitempty,gte=0"`
	Policyd_timestamp    int       `gorm:"type:int;default:null"`
	Smtpok               uint8     `gorm:"type:tinyint(1);not null;default:1"`
	Imapok               uint8     `gorm:"type:tinyint(1);not null;default:1"`
	Pop3ok               uint8     `gorm:"type:tinyint(1);not null;default:1"`
	Active               uint8     `gorm:"type:tinyint(1);not null;default:1"`
	Created              time.Time `gorm:"type:timestamp;not null;default:current_timestamp"`
	Transport            Transport `gorm:"foreignKey:Dominio"`
}
