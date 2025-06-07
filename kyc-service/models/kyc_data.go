package models

import (
	"time"
)

type KYCData struct {
	ID           int       `gorm:"primaryKey"`
	UserID       int       `gorm:"not null"`
	PhoneNumber  string    `gorm:"not null"`
	AadharNumber string    `gorm:"unique;not null"`
	PanNumber    string    `gorm:"unique;not null"`
	KYCStatus    string    `gorm:"default:'pending'"`
	CreatedAt    time.Time `gorm:"default:current_timestamp"`
}

func (KYCData) TableName() string {
	return "kyc_data"
}
