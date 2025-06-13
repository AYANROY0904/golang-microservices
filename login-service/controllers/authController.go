package controllers

import (
	"log"
	"time"

	"login-service/services"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
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

	// Return JWT and session ID directly (bypassing DB)
	c.JSON(200, gin.H{
		"jwt_token":  tokenString,
		"session_id": sessionID,
		"user_id":    0, // You can update this if you add user tracking later
		"expires_at": expiresAt,
		"message":    "OTP verified and session created successfully",
	})
}
