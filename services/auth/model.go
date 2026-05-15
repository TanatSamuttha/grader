package main

type User struct {
	UID string `gorm:"unique"`
	Google_UID string
	Email string
	Username string
}

type TokenDTO struct {
	Token string `json:"token"`
}