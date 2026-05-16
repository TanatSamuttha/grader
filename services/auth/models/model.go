package models

type User struct {
	UID string `gorm:"unique"`
	Google_UID string
	Email string
	Username string
	Version int8
}

type TokenDTO struct {
	Token string `json:"token"`
}