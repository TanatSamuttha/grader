package config

import (
	// "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB;

// func InitDatabase() error {
// 	var dns string = "https://fqbjupzsgdsyrmsqgtth.supabase.co"
// 	var err error;
// 	db, err = gorm.Open(postgres.Open(dns), &gorm.Config{})

// 	return err;
// }