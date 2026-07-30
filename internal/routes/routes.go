package routes

import (
	controller "github.com/rungrawee-popcorn/hospital-middleware-system/internal/controllers"
	middleware "github.com/rungrawee-popcorn/hospital-middleware-system/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	router *gin.Engine,
	staffController *controller.StaffController,
) {

	staffRoutes := router.Group("/staff")
	{

		staffRoutes.POST(
			"/create",
			staffController.CreateStaff,
		)

		staffRoutes.POST(
			"/login",
			staffController.Login,
		)

	}

	protected := router.Group("/api")

	protected.Use(
		middleware.JWTMiddleware(),
	)

	protected.GET(
		"/profile",
		func(ctx *gin.Context) {

			ctx.JSON(200, gin.H{
				"message": "protected route success",
			})

		},
	)

}