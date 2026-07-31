package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"github.com/rungrawee-popcorn/hospital-middleware-system/configs"
)

func JWTMiddleware() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" {

			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "missing authorization header",
			})

			ctx.Abort()
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {

			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "invalid authorization header",
			})

			ctx.Abort()
			return
		}

		tokenString := strings.TrimPrefix(
			authHeader,
			"Bearer ",
		)

		token, err := jwt.Parse(
			tokenString,
			func(token *jwt.Token) (interface{}, error) {

				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {

					return nil, jwt.ErrSignatureInvalid

				}

				return []byte(configs.Config.JWTSecret), nil

			},
		)

		if err != nil || !token.Valid {

			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "invalid token",
			})

			ctx.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok {

			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "invalid token claims",
			})

			ctx.Abort()
			return
		}

		staffID, ok := claims["staff_id"].(float64)

		if !ok {

			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "invalid staff id",
			})

			ctx.Abort()
			return
		}

		hospitalID, ok := claims["hospital_id"].(float64)

		if !ok {

			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "invalid hospital id",
			})

			ctx.Abort()
			return
		}

		username, ok := claims["username"].(string)

		if !ok {

			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "invalid username",
			})

			ctx.Abort()
			return
		}

		ctx.Set(
			"staff_id",
			uint(staffID),
		)

		ctx.Set(
			"hospital_id",
			uint(hospitalID),
		)

		ctx.Set(
			"username",
			username,
		)

		ctx.Next()

	}

}
