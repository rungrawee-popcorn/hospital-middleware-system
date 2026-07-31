package repository

import (
	model "github.com/rungrawee-popcorn/hospital-middleware-system/internal/models"

	"gorm.io/gorm"
)

type PatientRepository struct {
	DB *gorm.DB
}

func NewPatientRepository(
	db *gorm.DB,
) *PatientRepository {

	return &PatientRepository{
		DB: db,
	}
}

// SearchPatients searches patients by optional filters.
// HospitalID is mandatory to isolate hospital data.
func (r *PatientRepository) SearchPatients(
	hospitalID uint,
	filters map[string]interface{},
) ([]model.Patient, error) {

	var patients []model.Patient

	query := r.DB.
		Preload("Hospital").
		Where(
			"hospital_id = ?",
			hospitalID,
		)

	allowedFields := map[string]bool{

		"national_id": true,

		"passport_id": true,

		"first_name_th": true,

		"middle_name_th": true,

		"last_name_th": true,

		"date_of_birth": true,

		"phone_number": true,

		"email": true,
	}

	for field, value := range filters {

		if !allowedFields[field] {

			continue

		}

		if value != "" {

			query = query.Where(
				field+" = ?",
				value,
			)

		}

	}

	err := query.
		Find(&patients).
		Error

	if err != nil {

		return nil, err

	}

	return patients, nil
}

// FindByIdentifier searches patient by national ID or passport ID.
func (r *PatientRepository) FindByIdentifier(
	hospitalID uint,
	identifier string,
) (*model.Patient, error) {

	var patient model.Patient

	err := r.DB.
		Where(
			"hospital_id = ? AND (national_id = ? OR passport_id = ?)",
			hospitalID,
			identifier,
			identifier,
		).
		First(&patient).
		Error

	if err != nil {

		return nil, err

	}

	return &patient, nil
}

// CreatePatient saves patient information.
func (r *PatientRepository) CreatePatient(
	patient *model.Patient,
) error {

	return r.DB.Create(patient).Error
}
