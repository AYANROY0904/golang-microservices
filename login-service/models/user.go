package models

import (
	"time"


)

type User struct {
	ID          int       `gorm:"primaryKey"`
	PhoneNumber string    `gorm:"unique;not null"`
	CreatedAt   time.Time `gorm:"default:current_timestamp"`
}

func (User) TableName() string {
	return "users"
}
