package models

type Problems struct {
	ID				uint32	`gorm:"primaryKey"`
	ProblemID		string	`gorm:"unique"`
	Name			string
	AuthorUID		string
	TestCasesSize	uint8
	TimeLimit		uint8
	MemoryLimit		uint8
}