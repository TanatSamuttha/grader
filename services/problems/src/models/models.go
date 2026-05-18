package models

type Problems struct {
	ID				uint32	`gorm:"primaryKey"`
	FileID			string	`gorm:"unique"`
	TestCasesSize	uint8
	TimeLimit		int8
	MemoryLimit		int8
}