package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateStaffSuccess(t *testing.T) {
	setupTestDatabase(t)
	clearDatabase()

	router := setupRouter()

	hospital := createTestHospital()

	body := map[string]interface{}{
		"username":    "doctor01",
		"password":    "123456",
		"hospital_id": hospital.ID,
	}

	jsonBody, _ := json.Marshal(body)

	req := httptest.NewRequest(
		http.MethodPost,
		"/staff/create",
		bytes.NewBuffer(jsonBody),
	)

	req.Header.Set(
		"Content-Type",
		"application/json",
	)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(
		t,
		http.StatusCreated,
		rec.Code,
	)
}

func TestCreateStaffDuplicate(t *testing.T) {
	setupTestDatabase(t)
	clearDatabase()

	router := setupRouter()

	hospital := createTestHospital()

	jsonBody := []byte(`{
		"username":"doctor01",
		"password":"123456",
		"hospital_id":1
	}`)

	for i := 0; i < 2; i++ {
		req := httptest.NewRequest(
			http.MethodPost,
			"/staff/create",
			bytes.NewBuffer(jsonBody),
		)

		req.Header.Set(
			"Content-Type",
			"application/json",
		)

		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		if i == 1 {
			assert.Equal(
				t,
				http.StatusConflict,
				rec.Code,
			)
		}
	}

	_ = hospital
}

func TestStaffLoginSuccess(t *testing.T) {
	setupTestDatabase(t)
	clearDatabase()

	router := setupRouter()

	hospital := createTestHospital()

	createBody := []byte(`{
		"username":"doctor01",
		"password":"123456",
		"hospital_id":1
	}`)

	req := httptest.NewRequest(
		http.MethodPost,
		"/staff/create",
		bytes.NewBuffer(createBody),
	)

	req.Header.Set(
		"Content-Type",
		"application/json",
	)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	loginBody := []byte(`{
		"username":"doctor01",
		"password":"123456",
		"hospital_id":1
	}`)

	req = httptest.NewRequest(
		http.MethodPost,
		"/staff/login",
		bytes.NewBuffer(loginBody),
	)

	req.Header.Set(
		"Content-Type",
		"application/json",
	)

	rec = httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(
		t,
		http.StatusOK,
		rec.Code,
	)

	_ = hospital
}

func TestStaffLoginWrongPassword(t *testing.T) {
	setupTestDatabase(t)
	clearDatabase()

	router := setupRouter()

	createTestHospital()

	createBody := []byte(`{
		"username":"doctor01",
		"password":"123456",
		"hospital_id":1
	}`)

	req := httptest.NewRequest(
		http.MethodPost,
		"/staff/create",
		bytes.NewBuffer(createBody),
	)

	req.Header.Set(
		"Content-Type",
		"application/json",
	)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	loginBody := []byte(`{
		"username":"doctor01",
		"password":"wrong",
		"hospital_id":1
	}`)

	req = httptest.NewRequest(
		http.MethodPost,
		"/staff/login",
		bytes.NewBuffer(loginBody),
	)

	req.Header.Set(
		"Content-Type",
		"application/json",
	)

	rec = httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(
		t,
		http.StatusUnauthorized,
		rec.Code,
	)
}