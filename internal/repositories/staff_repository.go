package repository

import (
	model "github.com/rungrawee-popcorn/hospital-middleware-system/internal/models"

	"gorm.io/gorm"
)

type StaffRepository struct {
	DB *gorm.DB
}

func NewStaffRepository(db *gorm.DB) *StaffRepository {
	return &StaffRepository{
		DB: db,
	}
}

// CreateStaff inserts a new staff record into the database.
func (r *StaffRepository) CreateStaff(staff *model.Staff) error {
	return r.DB.Create(staff).Error
}

// FindByUsernameAndHospital finds a staff by username and hospital.
func (r *StaffRepository) FindByUsernameAndHospital(
	username string,
	hospitalID uint,
) (*model.Staff, error) {

	var staff model.Staff

	err := r.DB.
		Preload("Hospital").
		Where(
			"username = ? AND hospital_id = ?",
			username,
			hospitalID,
		).
		First(&staff).Error

	if err != nil {
		return nil, err
	}

	return &staff, nil
}

// UsernameExistsInHospital checks whether the username already exists
// within the same hospital.
func (r *StaffRepository) UsernameExistsInHospital(
	username string,
	hospitalID uint,
) bool {

	var count int64

	r.DB.
		Model(&model.Staff{}).
		Where(
			"username = ? AND hospital_id = ?",
			username,
			hospitalID,
		).
		Count(&count)

	return count > 0
}

// FindByID returns a staff by primary key.
func (r *StaffRepository) FindByID(
	id uint,
) (*model.Staff, error) {

	var staff model.Staff

	err := r.DB.
		Preload("Hospital").
		First(&staff, id).Error

	if err != nil {
		return nil, err
	}

	return &staff, nil
}
