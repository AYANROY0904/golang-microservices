package controllers

import (
	"shared/utils"
	"user-profile-service/services"

	"net/http"
	"github.com/gin-gonic/gin"
)


func HandleProfile(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
		return
	}

	claims, err := utils.ValidateJWT(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	phoneNumber, ok := claims["phone_number"].(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID in token"})
		return
	}

	profile, err := services.FetchUserProfileWithKYC(phoneNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user profile"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"profile": profile})
}
