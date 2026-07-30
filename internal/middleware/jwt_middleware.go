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

		tokenString := strings.Replace(
			authHeader,
			"Bearer ",
			"",
			1,
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

		ctx.Next()

	}

}