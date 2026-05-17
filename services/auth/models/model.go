package models

type User struct {
	ID			int8 `gorm:"primaryKey"`
	UID 		string `gorm:"unique"`
	Google_UID 	string
	Email 		string
	Username 	string
	PhotoURL 	string
	Version 	int8
}

type TokenDTO struct {
	Token string `json:"token"`
}

type UserDTO struct {
	Username string `json:"username"`
	PhotoURL string `json:"photoURL"`
}