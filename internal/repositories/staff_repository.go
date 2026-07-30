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

func (r *StaffRepository) CreateStaff(staff *model.Staff) error {
	return r.DB.Create(staff).Error
}

func (r *StaffRepository) FindByUsername(username string) (*model.Staff, error) {

	var staff model.Staff

	err := r.DB.
		Preload("Hospital").
		Where("username = ?", username).
		First(&staff).Error

	if err != nil {
		return nil, err
	}

	return &staff, nil
}

func (r *StaffRepository) UsernameExists(username string) bool {

	var count int64

	r.DB.Model(&model.Staff{}).
		Where("username = ?", username).
		Count(&count)

	return count > 0
}