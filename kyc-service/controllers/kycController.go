
package controllers

import (
	"net/http"
	"kyc-service/services"
	"shared/utils"
	"github.com/gin-gonic/gin"
	"github.com/getsentry/sentry-go"
)

type KYCData struct {
	AadharNumber string `json:"aadhar_number"`
	PanNumber    string `json:"pan_number"`
}

func HandleKYC(c *gin.Context) {
	var request KYCData

	if err := c.ShouldBindJSON(&request); err != nil {
		sentry.CaptureException(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
		return
	}

	claims, err := utils.ValidateJWT(tokenString)
	if err != nil {
		sentry.CaptureException(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		return
	}

	phoneNumber, ok := claims["phone_number"].(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID in token"})
		return
	}

	message, err := services.StoreKYCData( phoneNumber, request.AadharNumber, request.PanNumber)
	if err != nil {
		sentry.CaptureException(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process KYC data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": message})
}
