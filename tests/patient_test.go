package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	model "github.com/rungrawee-popcorn/hospital-middleware-system/internal/models"
)

func TestSearchPatientSuccess(t *testing.T) {
	setupTestDatabase(t)
	clearDatabase()

	router := setupRouter()

	hospital := createTestHospital()

	patient := model.Patient{
		HospitalID:  hospital.ID,
		NationalID:  "123456789",
		FirstNameTH: "Somchai",
		LastNameTH:  "Test",
		DateOfBirth: time.Now(),
	}

	testDB.Create(&patient)

	token := generateTestToken(
		1,
		hospital.ID,
		"doctor01",
	)

	req := httptest.NewRequest(
		http.MethodGet,
		"/api/patient/search?national_id=123456789",
		nil,
	)

	req.Header.Set(
		"Authorization",
		"Bearer "+token,
	)

	rec := httptest.NewRecorder()

	router.ServeHTTP(
		rec,
		req,
	)

	assert.Equal(
		t,
		http.StatusOK,
		rec.Code,
	)
}

func TestSearchPatientWithoutToken(t *testing.T) {
	setupTestDatabase(t)
	clearDatabase()

	router := setupRouter()

	req := httptest.NewRequest(
		http.MethodGet,
		"/api/patient/search?national_id=123456789",
		nil,
	)

	rec := httptest.NewRecorder()

	router.ServeHTTP(
		rec,
		req,
	)

	assert.Equal(
		t,
		http.StatusUnauthorized,
		rec.Code,
	)
}

func TestSearchPatientWrongHospital(t *testing.T) {
	setupTestDatabase(t)
	clearDatabase()

	router := setupRouter()

	hospital := createTestHospital()

	token := generateTestToken(
		1,
		hospital.ID+99,
		"doctor01",
	)

	req := httptest.NewRequest(
		http.MethodGet,
		"/api/patient/search?national_id=123456789",
		nil,
	)

	req.Header.Set(
		"Authorization",
		"Bearer "+token,
	)

	rec := httptest.NewRecorder()

	router.ServeHTTP(
		rec,
		req,
	)

	assert.Equal(
		t,
		http.StatusOK,
		rec.Code,
	)
}