package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)


func GenerateToken(
	staffID uint,
	username string,
	hospitalID uint,
	secret string,
) (string, error) {


	claims := jwt.MapClaims{

		"staff_id": staffID,

		"username": username,

		"hospital_id": hospitalID,

		"exp": time.Now().
			Add(time.Hour * 24).
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