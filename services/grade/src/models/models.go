package models

import (
	"time"

	"github.com/gofiber/contrib/v3/websocket"
)

type Problem struct {
	ID            uint32 `gorm:"primaryKey"`
	ProblemID     string `gorm:"unique"`
	Name          string
	AuthorUID     string
	TestCasesSize uint8
	TimeLimit     uint64
	MemoryLimit   uint16
	Visibility    string
}

type User struct {
	ID         uint32 `gorm:"primaryKey"`
	UID        string `gorm:"unique"`
	Google_UID string
	Email      string
	Username   string
	PhotoURL   string
	Role       string
	Version    uint16
}

type Job struct {
	ID        string
	UID       string
	ProblemID string
	Code      string
	Lang      string
}

type CodeDTO struct {
	Code string `json:"code"`
	Lang string `json:"lang"`
}

type GradeResJob struct {
	JobID  	string
	Task   	int
	Score	bool
	Compile	bool
	Error  	string
	Conn   	*websocket.Conn
}

type GradeResDTO struct {
	Task   	int
	Score 	bool
	Compile bool
	Error  	string
}

type Submission struct {
	ID			int		`gorm:"primaryKey"`
	UID			string
	Score		int
	Error		string
	CreatedAt	time.Time
}