package models

import (
	"log"
	"time"

	"shared/utils" // Update the import path based on your module
)

type Session struct {
	ID          int       `gorm:"primaryKey;autoIncrement"`
	UserID      int       `gorm:"not null"`
	PhoneNumber string    `gorm:"not null"`
	SessionID   string    `gorm:"unique;not null"`
	JWTToken    string    `gorm:"not null"`
	ExpiresAt   time.Time `gorm:"not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}

func (Session) TableName() string {
	return "sessions"
}

func MigrateSession() {
	err := utils.DB.AutoMigrate(&Session{})
	if err != nil {
		log.Fatalf("Session table migration failed: %v", err)
	}
}
