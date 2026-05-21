package models

type Problem struct {
	ID				uint32				`gorm:"primaryKey"`
	ProblemID		string				`gorm:"unique"`
	Name			string
	AuthorUID		string
	TestCasesSize	uint8
	TimeLimit		uint8
	MemoryLimit		uint8
	Visibility		string
}

type User struct {
	ID				uint32 				`gorm:"primaryKey"`
	UID 			string 				`gorm:"unique"`
	Google_UID 		string
	Email 			string
	Username 		string
	PhotoURL 		string
	Role			string
	Version 		uint16
}

type Job struct {
	ID				string
	UID				string
	ProblemID		string
	Code			string
	Lang			string
}

type CodeDTO struct {
	Code			string				`json:"code"`
	Lang			string				`json:"lang"`
}