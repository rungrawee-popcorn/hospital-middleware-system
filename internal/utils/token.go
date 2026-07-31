package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateToken creates a JWT token for an authenticated staff.
func GenerateToken(
	staffID uint,
	username string,
	hospitalID uint,
	secret string,
) (string, error) {

	claims := jwt.MapClaims{
		"staff_id":    staffID,
		"username":    username,
		"hospital_id": hospitalID,
		"exp": time.Now().
			Add(24 * time.Hour).
			Unix(),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString(
		[]byte(secret),
	)
}
