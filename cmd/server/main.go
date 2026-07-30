package main

import (
	"github.com/gin-gonic/gin"

	"github.com/rungrawee-popcorn/hospital-middleware-system/configs"

	controller "github.com/rungrawee-popcorn/hospital-middleware-system/internal/controllers"

	repository "github.com/rungrawee-popcorn/hospital-middleware-system/internal/repositories"

	service "github.com/rungrawee-popcorn/hospital-middleware-system/internal/services"

	route "github.com/rungrawee-popcorn/hospital-middleware-system/internal/routes"
)

func main() {

	// Load environment variables

	configs.LoadConfig()

	// Connect database

	configs.ConnectDatabase()

	// Run migration

	configs.AutoMigrate()

	// Dependency Injection

	staffRepository := repository.NewStaffRepository(
		configs.DB,
	)

	staffService := service.NewStaffService(
		staffRepository,
	)

	staffController := controller.NewStaffController(
		staffService,
	)

	// Gin Router

	router := gin.Default()

	route.SetupRoutes(
		router,
		staffController,
	)

	router.Run(":8080")

}