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

func (s *StaffService) CreateStaff(
	staff *model.Staff,
) error {

	if s.Repository.UsernameExists(staff.Username) {
		return errors.New("username already exists")
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

func (s *StaffService) Login(
	username string,
	password string,
	secret string,
) (*model.Staff, string, error) {

	staff, err := s.Repository.FindByUsername(username)

	if err != nil {
		return nil, "", errors.New("invalid username or password")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(staff.Password),
		[]byte(password),
	)

	if err != nil {
		return nil, "", errors.New("invalid username or password")
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