package models

type Problems struct {
	ID				uint32	`gorm:"primaryKey"`
	ProblemID		string	`gorm:"unique"`
	TestCasesSize	uint8
	TimeLimit		uint8
	MemoryLimit		uint8
}