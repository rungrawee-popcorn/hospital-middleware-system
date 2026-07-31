package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	service "github.com/rungrawee-popcorn/hospital-middleware-system/internal/services"
)

type PatientController struct {
	PatientService *service.PatientService
}

func NewPatientController(
	patientService *service.PatientService,
) *PatientController {

	return &PatientController{
		PatientService: patientService,
	}
}

type PatientSearchRequest struct {
	NationalID string `form:"national_id"`

	PassportID string `form:"passport_id"`

	FirstName string `form:"first_name"`

	MiddleName string `form:"middle_name"`

	LastName string `form:"last_name"`

	DateOfBirth string `form:"date_of_birth"`

	PhoneNumber string `form:"phone_number"`

	Email string `form:"email"`
}

// GET /api/patient/search
func (c *PatientController) SearchPatient(
	ctx *gin.Context,
) {

	hospitalValue, exists := ctx.Get(
		"hospital_id",
	)

	if !exists {

		ctx.JSON(
			http.StatusUnauthorized,
			gin.H{
				"message": "hospital information missing",
			},
		)

		return
	}

	hospitalID, ok := hospitalValue.(uint)

	if !ok {

		ctx.JSON(
			http.StatusUnauthorized,
			gin.H{
				"message": "invalid hospital information",
			},
		)

		return
	}

	var request PatientSearchRequest

	if err := ctx.ShouldBindQuery(
		&request,
	); err != nil {

		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": "invalid request",
				"error":   err.Error(),
			},
		)

		return
	}

	filters := map[string]interface{}{}

	if request.NationalID != "" {

		filters["national_id"] = request.NationalID

	}

	if request.PassportID != "" {

		filters["passport_id"] = request.PassportID

	}

	if request.FirstName != "" {

		filters["first_name_th"] = request.FirstName

	}

	if request.MiddleName != "" {

		filters["middle_name_th"] = request.MiddleName

	}

	if request.LastName != "" {

		filters["last_name_th"] = request.LastName

	}

	if request.DateOfBirth != "" {

		filters["date_of_birth"] = request.DateOfBirth

	}

	if request.PhoneNumber != "" {

		filters["phone_number"] = request.PhoneNumber

	}

	if request.Email != "" {

		filters["email"] = request.Email

	}

	patients, err := c.PatientService.SearchPatient(
		hospitalID,
		filters,
	)

	if err != nil {

		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"message": err.Error(),
			},
		)

		return
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"message": "patients found",
			"data":    patients,
		},
	)
}
