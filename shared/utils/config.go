package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost       	   string
	DBUser       	   string
	DBPassword   	   string
	DBName       	   string
	DBPort       	   string
	JwtSecret    	   string
	TWILIO_ACCOUNT_SID string
	TWILIO_AUTH_TOKEN  string
	TWILIO_SERVICE_SID string
	SENTRY_DSN         string
}

func LoadConfig() Config {
	err := godotenv.Load("../shared/.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err) 
	}

	return Config{
		DBHost:       os.Getenv("DB_HOST"),
		DBUser:       os.Getenv("DB_USER"),
		DBPassword:   os.Getenv("DB_PASSWORD"),
		DBName:       os.Getenv("DB_NAME"),
		DBPort:       os.Getenv("DB_PORT"),
		JwtSecret:    os.Getenv("JWT_SECRET"),
		TWILIO_ACCOUNT_SID: os.Getenv("TWILIO_ACCOUNT_SID"),
		TWILIO_AUTH_TOKEN:  os.Getenv("TWILIO_AUTH_TOKEN"),
		TWILIO_SERVICE_SID: os.Getenv("TWILIO_SERVICE_SID"),
		SENTRY_DSN: os.Getenv("SENTRY_DSN"),
	}
}
