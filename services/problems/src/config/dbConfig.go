package config

import (
	"os"
	"problems/models"

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
		return err;
	}

	DB.AutoMigrate(&models.Problems{});

	return nil;
}