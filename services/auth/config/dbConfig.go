package config

import (
	"auth/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB;

func InitDatabase() error {
	var dns string = os.Getenv("DATABASE_URL");
	var err error
	DB, err = gorm.Open(postgres.Open(dns), &gorm.Config{});
	if err != nil {
		return err;
	}

	DB.AutoMigrate(&models.User{});

	return nil;
}