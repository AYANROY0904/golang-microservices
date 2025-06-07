package controllers

import (
	"log"
	"time"
	"login-service/services"
	"shared/utils"
	"github.com/gin-gonic/gin"
	"github.com/getsentry/sentry-go"
)

func Login(c *gin.Context) {
	var request struct {
		PhoneNumber string `json:"phone_number"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		sentry.WithScope(func(scope *sentry.Scope) {
			scope.SetTag("service", "login-service")
			sentry.CaptureException(err)
		})
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	err := services.SendOTP(request.PhoneNumber)
	if err != nil {
		sentry.WithScope(func(scope *sentry.Scope) {
			scope.SetTag("phone_number", request.PhoneNumber)
			sentry.CaptureException(err)
		})
		c.JSON(500, gin.H{"error": "Failed to send OTP"})
		return
	}
	c.JSON(200, gin.H{"message": "OTP sent successfully"})
}

func VerifyOTP(c *gin.Context) {
	var request struct {
		PhoneNumber string `json:"phone_number"`
		OTP         string `json:"otp"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		sentry.WithScope(func(scope *sentry.Scope) {
			scope.SetTag("service", "login-service")
			sentry.CaptureException(err)
		})
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	isValid, err := services.VerifyOTP(request.PhoneNumber, request.OTP)
	if err != nil || !isValid {
		sentry.WithScope(func(scope *sentry.Scope) {
			scope.SetTag("phone_number", request.PhoneNumber)
			sentry.CaptureException(err)
		})
		c.JSON(400, gin.H{"error": "Invalid OTP"})
		return
	}


	// Generate JWT and session ID
	tokenString, sessionID, err := services.GenerateJWT(0, request.PhoneNumber)
	if err != nil {
		sentry.WithScope(func(scope *sentry.Scope) {
			scope.SetTag("phone_number", request.PhoneNumber)
			sentry.CaptureException(err)
		})
		log.Println("Error generating JWT:", err)
		c.JSON(500, gin.H{"error": "Failed to generate JWT"})
		return
	}

	expiresAt := time.Now().Add(30 * time.Minute)

	// Call the stored procedure
	var userID int
	var message string
	result := utils.DB.Raw(`
		SELECT * FROM verify_otp_and_create_session(?, ?, ?, ?)
	`, request.PhoneNumber, tokenString, sessionID, expiresAt).Row()

	if err := result.Scan(&userID, &message); err != nil {
		sentry.WithScope(func(scope *sentry.Scope) {
			scope.SetTag("phone_number", request.PhoneNumber)
			sentry.CaptureException(err)
		})
		log.Println("Error executing stored procedure:", err)
		c.JSON(500, gin.H{"error": "Failed to verify OTP and create session"})
		return
	}

	c.JSON(200, gin.H{
		"jwt_token":  tokenString,
		"session_id": sessionID,
		"user_id": userID,
		"message": message,
	})
}
