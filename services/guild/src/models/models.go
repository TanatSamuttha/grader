package models

type Guild struct {
	ID			uint32	`gorm"primaryKey"`
	GuildID		string	`gorm:"unique"`
	LeaderID	string
}

type GuildProblem struct {
	ID			uint32	`gorm:"primaryKey"`
	ProblemID	string	`gorm:"unique"`
	GuildID		string
}

type GuildMember struct {
	ID			uint32 	`gorm:"primaryKey"`
	UID 		string 	`gorm:"unique"`
	GuildID		string
}

type User struct {
	ID			uint32 	`gorm:"primaryKey"`
	UID 		string 	`gorm:"unique"`
	Google_UID 	string
	Email 		string
	Username 	string
	PhotoURL 	string
	Role		string
	Version 	uint16
}