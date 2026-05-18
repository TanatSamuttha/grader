package models

type Problem struct {
	ID				uint32	`gorm:"primaryKey"`
	ProblemID		string	`gorm:"unique"`
	Name			string
	AuthorUID		string
	TestCasesSize	uint8
	TimeLimit		uint8
	MemoryLimit		uint8
}

type User struct {
	ID				uint32 	`gorm:"primaryKey"`
	UID 			string 	`gorm:"unique"`
	Google_UID 		string
	Email 			string
	Username 		string
	PhotoURL 		string
	Role			string
	Version 		uint16
}

type ProblemDTO struct {
	Name			string	`json:"name"`
	TimeLimit		uint8	`json:"timeLimit"`
	MemoryLimit		uint8	`json:"memoryLimit"`
}