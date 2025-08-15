package config

import (
	"practice_account/pkg/domain"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() (*gorm.DB, error) {

	var err error
	dsn := "web:password@/email?parseTime=true"

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = DB.AutoMigrate(&domain.Aliases{}, &domain.Transport{}, &domain.Users{})
	if err != nil {
		return nil, err
	}

	return DB, nil
}

