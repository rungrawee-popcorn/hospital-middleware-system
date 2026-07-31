package main

import (
	"github.com/gin-gonic/gin"

	"github.com/rungrawee-popcorn/hospital-middleware-system/configs"

	controller "github.com/rungrawee-popcorn/hospital-middleware-system/internal/controllers"

	hospital "github.com/rungrawee-popcorn/hospital-middleware-system/internal/hospital"

	repository "github.com/rungrawee-popcorn/hospital-middleware-system/internal/repositories"

	service "github.com/rungrawee-popcorn/hospital-middleware-system/internal/services"

	route "github.com/rungrawee-popcorn/hospital-middleware-system/internal/routes"
)

func main() {

	// Load environment variables
	configs.LoadConfig()

	// Connect database
	configs.ConnectDatabase()

	// Run database migration
	configs.AutoMigrate()

	// =========================
	// Staff Module
	// =========================

	staffRepository := repository.NewStaffRepository(
		configs.DB,
	)

	staffService := service.NewStaffService(
		staffRepository,
	)

	staffController := controller.NewStaffController(
		staffService,
	)

	// =========================
	// Patient Module
	// =========================

	patientRepository := repository.NewPatientRepository(
		configs.DB,
	)

	hospitalClient := hospital.NewHospitalClient(
		configs.Config.HospitalAAPIURL,
	)

	patientService := service.NewPatientService(
		patientRepository,
		hospitalClient,
	)

	patientController := controller.NewPatientController(
		patientService,
	)

	// =========================
	// Gin Router
	// =========================

	router := gin.Default()

	// Security: do not trust all proxies
	router.SetTrustedProxies(nil)

	route.SetupRoutes(
		router,
		staffController,
		patientController,
	)

	router.Run(":8080")
}
