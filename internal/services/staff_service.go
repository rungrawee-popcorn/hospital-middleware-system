package service

import (
	"errors"

	model "github.com/rungrawee-popcorn/hospital-middleware-system/internal/models"
	repository "github.com/rungrawee-popcorn/hospital-middleware-system/internal/repositories"
	utils "github.com/rungrawee-popcorn/hospital-middleware-system/internal/utils"

	"golang.org/x/crypto/bcrypt"
)

type StaffService struct {
	Repository *repository.StaffRepository
}

func NewStaffService(
	repo *repository.StaffRepository,
) *StaffService {

	return &StaffService{
		Repository: repo,
	}
}

// CreateStaff creates a new hospital staff account.
func (s *StaffService) CreateStaff(
	staff *model.Staff,
) error {

	if s.Repository.UsernameExistsInHospital(
		staff.Username,
		staff.HospitalID,
	) {
		return errors.New("username already exists in this hospital")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(staff.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	staff.Password = string(hashedPassword)

	return s.Repository.CreateStaff(staff)
}

// Login authenticates a staff account and generates a JWT.
func (s *StaffService) Login(
	username string,
	password string,
	hospitalID uint,
	secret string,
) (*model.Staff, string, error) {

	staff, err := s.Repository.FindByUsernameAndHospital(
		username,
		hospitalID,
	)

	if err != nil {
		return nil, "", errors.New("invalid username, password or hospital")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(staff.Password),
		[]byte(password),
	)

	if err != nil {
		return nil, "", errors.New("invalid username, password or hospital")
	}

	token, err := utils.GenerateToken(
		staff.ID,
		staff.Username,
		staff.HospitalID,
		secret,
	)

	if err != nil {
		return nil, "", err
	}

	return staff, token, nil
}

// GetStaffByID returns a staff by id.
func (s *StaffService) GetStaffByID(
	id uint,
) (*model.Staff, error) {

	return s.Repository.FindByID(id)
}
