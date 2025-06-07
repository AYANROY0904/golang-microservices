package services

import (
	"log"
	"time"
	"shared/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/getsentry/sentry-go"
)


func GenerateJWT(userID int, phoneNumber string) (string, string, error) {

	sessionID := uuid.New().String()
	expiresAt := time.Now().Add(30 * time.Minute)

	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":    userID,
		"phone_number": phoneNumber,
		"session_id": sessionID,
		"exp":        expiresAt.Unix(),
	})

	// Sign the token
	tokenString, err := token.SignedString([]byte(utils.LoadConfig().JwtSecret))
	if err != nil {
		sentry.WithScope(func(scope *sentry.Scope) {
			scope.SetTag("phone_number", phoneNumber)
			sentry.CaptureException(err)
		})
		log.Println("Error generating JWT:", err)
		return "", "", err
	}

	return tokenString, sessionID, nil
}

