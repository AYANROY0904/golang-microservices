package models

import (
	"log"
	"time"

	"shared/utils" // Update the import path based on your module
)

type KYCData struct {
	ID           int       `gorm:"primaryKey;autoIncrement"`
	UserID       int       `gorm:"not null"`
	PhoneNumber  string    `gorm:"not null"`
	AadharNumber string    `gorm:"unique;not null"`
	PanNumber    string    `gorm:"unique;not null"`
	KYCStatus    string    `gorm:"default:'pending'"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
}

func (KYCData) TableName() string {
	return "kyc_data"
}

func MigrateKYC() {
	err := utils.DB.AutoMigrate(&KYCData{})
	if err != nil {
		log.Fatalf("KYCData table migration failed: %v", err)
	}
}
