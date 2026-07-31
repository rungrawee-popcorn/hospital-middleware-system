package tests

import (
	"fmt"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"

	"github.com/rungrawee-popcorn/hospital-middleware-system/configs"
	controller "github.com/rungrawee-popcorn/hospital-middleware-system/internal/controllers"
	hospital "github.com/rungrawee-popcorn/hospital-middleware-system/internal/hospital"
	middlewares "github.com/rungrawee-popcorn/hospital-middleware-system/internal/middleware"
	model "github.com/rungrawee-popcorn/hospital-middleware-system/internal/models"
	repository "github.com/rungrawee-popcorn/hospital-middleware-system/internal/repositories"
	service "github.com/rungrawee-popcorn/hospital-middleware-system/internal/services"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var testDB *gorm.DB

func setupTestDatabase(t *testing.T) {

	err := godotenv.Load("../.env.test")

	require.NoError(
		t,
		err,
		"cannot load .env.test",
	)

	configs.LoadConfig()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		configs.Config.DBHost,
		configs.Config.DBUser,
		configs.Config.DBPassword,
		configs.Config.DBName,
		configs.Config.DBPort,
	)

	db, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{},
	)

	require.NoError(
		t,
		err,
	)

	testDB = db

	err = testDB.AutoMigrate(
		&model.Hospital{},
		&model.Staff{},
		&model.Patient{},
	)

	require.NoError(
		t,
		err,
	)
}

func clearDatabase() {

	testDB.Exec(
		"DELETE FROM patients",
	)

	testDB.Exec(
		"DELETE FROM staffs",
	)

	testDB.Exec(
		"DELETE FROM hospitals",
	)

	testDB.Exec(
		"ALTER SEQUENCE hospitals_id_seq RESTART WITH 1",
	)

	testDB.Exec(
		"ALTER SEQUENCE staffs_id_seq RESTART WITH 1",
	)

	testDB.Exec(
		"ALTER SEQUENCE patients_id_seq RESTART WITH 1",
	)
}

func createTestHospital() model.Hospital {

	hospital := model.Hospital{
		Name: "Hospital A",
	}

	testDB.Create(
		&hospital,
	)

	return hospital
}

func setupRouter() *gin.Engine {

	staffRepo := repository.NewStaffRepository(
		testDB,
	)

	staffService := service.NewStaffService(
		staffRepo,
	)

	staffController := controller.NewStaffController(
		staffService,
	)

	patientRepo := repository.NewPatientRepository(
		testDB,
	)

	hospitalClient := hospital.NewHospitalClient(
		"",
	)

	patientService := service.NewPatientService(
		patientRepo,
		hospitalClient,
	)

	patientController := controller.NewPatientController(
		patientService,
	)

	router := gin.Default()

	router.POST(
		"/staff/create",
		staffController.CreateStaff,
	)

	router.POST(
		"/staff/login",
		staffController.Login,
	)

	patientRoute := router.Group(
		"/api/patient",
	)

	patientRoute.Use(
		middlewares.JWTMiddleware(),
	)

	patientRoute.GET(
		"/search",
		patientController.SearchPatient,
	)

	return router
}

func generateTestToken(
	staffID uint,
	hospitalID uint,
	username string,
) string {

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"staff_id":    staffID,
			"hospital_id": hospitalID,
			"username":    username,
		},
	)

	tokenString, err := token.SignedString(
		[]byte(configs.Config.JWTSecret),
	)

	if err != nil {
		return ""
	}

	return tokenString
}