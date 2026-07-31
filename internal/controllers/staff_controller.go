package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/rungrawee-popcorn/hospital-middleware-system/configs"

	model "github.com/rungrawee-popcorn/hospital-middleware-system/internal/models"
	service "github.com/rungrawee-popcorn/hospital-middleware-system/internal/services"
)

type StaffController struct {
	StaffService *service.StaffService
}

func NewStaffController(
	staffService *service.StaffService,
) *StaffController {

	return &StaffController{
		StaffService: staffService,
	}
}

type CreateStaffRequest struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	HospitalID uint   `json:"hospital_id"`
}

type LoginRequest struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	HospitalID uint   `json:"hospital_id"`
}

// POST /staff/create
func (c *StaffController) CreateStaff(ctx *gin.Context) {

	var request CreateStaffRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
			"error":   err.Error(),
		})

		return
	}

	if request.Username == "" ||
		request.Password == "" ||
		request.HospitalID == 0 {

		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "username, password and hospital_id are required",
		})

		return
	}

	staff := model.Staff{
		Username:   request.Username,
		Password:   request.Password,
		HospitalID: request.HospitalID,
	}

	err := c.StaffService.CreateStaff(&staff)

	if err != nil {

		ctx.JSON(http.StatusConflict, gin.H{
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "staff created successfully",
		"staff": gin.H{
			"id":          staff.ID,
			"username":    staff.Username,
			"hospital_id": staff.HospitalID,
		},
	})
}

// POST /staff/login
func (c *StaffController) Login(ctx *gin.Context) {

	var request LoginRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
			"error":   err.Error(),
		})

		return
	}

	if request.Username == "" ||
		request.Password == "" ||
		request.HospitalID == 0 {

		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "username, password and hospital_id are required",
		})

		return
	}

	staff, token, err := c.StaffService.Login(
		request.Username,
		request.Password,
		request.HospitalID,
		configs.Config.JWTSecret,
	)

	if err != nil {

		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "login successful",
		"token":   token,
		"staff": gin.H{
			"id":          staff.ID,
			"username":    staff.Username,
			"hospital_id": staff.HospitalID,
			"hospital":    staff.Hospital.Name,
		},
	})
}
