package repository

import (
	model "github.com/rungrawee-popcorn/hospital-middleware-system/internal/models"

	"gorm.io/gorm"
)

type HospitalRepository struct {
	DB *gorm.DB
}

func NewHospitalRepository(
	db *gorm.DB,
) *HospitalRepository {

	return &HospitalRepository{
		DB: db,
	}
}

func (r *HospitalRepository) FindByID(
	id uint,
) (*model.Hospital,error){

	var hospital model.Hospital

	err := r.DB.
		First(&hospital,id).
		Error


	if err != nil {
		return nil,err
	}

	return &hospital,nil
}