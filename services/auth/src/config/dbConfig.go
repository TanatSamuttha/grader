package config

import (
	"auth/models"
	"errors"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB;

func InitDatabase() error {
	var dsn string = os.Getenv("DATABASE_URL");
	var err error;
	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{});
	if err != nil {
		return errors.New("Error connect to data base -> " + err.Error());
	}

	DB.AutoMigrate(&models.User{});

	return nil;
}