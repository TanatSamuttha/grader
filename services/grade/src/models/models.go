package models

import "github.com/gofiber/contrib/v3/websocket"

type Problem struct {
	ID            uint32 `gorm:"primaryKey"`
	ProblemID     string `gorm:"unique"`
	Name          string
	AuthorUID     string
	TestCasesSize uint8
	TimeLimit     uint8
	MemoryLimit   uint8
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
	JobID  string
	Task   int
	Result bool
	Error  string
	Conn   *websocket.Conn
}

type GradeResDTO struct {
	Task   int
	Result bool
	Error  string
}