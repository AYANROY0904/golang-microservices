package models

import (
	"log"
	"time"

	"shared/utils" // Update the import path based on your module
)

type User struct {
	ID          int       `gorm:"primaryKey;autoIncrement"`
	PhoneNumber string    `gorm:"unique;not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}

func (User) TableName() string {
	return "users"
}

func MigrateUser() {
	err := utils.DB.AutoMigrate(&User{})
	if err != nil {
		log.Fatalf("User table migration failed: %v", err)
	}
}
