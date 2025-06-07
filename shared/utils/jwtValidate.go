package utils

import (
	"fmt"
	"strings"
	"github.com/golang-jwt/jwt/v5"
	"log"
)

func ValidateJWT(tokenString string) (jwt.MapClaims, error) {
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(LoadConfig().JwtSecret), nil
	})
	if err != nil || !token.Valid {
		log.Println("Invalid or expired JWT:", err)
		return nil, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("could not extract claims")
	}

	return claims, nil
}
