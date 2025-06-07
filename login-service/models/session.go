package models

import (
	"time"
)

type Session struct {
	ID         int       `gorm:"primaryKey"`
	UserID     int       `gorm:"not null"`
	PhoneNumber string   `gorm:"not null"`
	SessionID  string `gorm:"unique;not null"`
	JWTToken   string    `gorm:"not null"`
	ExpiresAt  time.Time `gorm:"not null"`
	CreatedAt  time.Time `gorm:"default:current_timestamp"`
}

func (Session) TableName() string {
	return "sessions"
}
