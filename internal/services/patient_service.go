package service

import (
	"errors"

	model "github.com/rungrawee-popcorn/hospital-middleware-system/internal/models"

	hospital "github.com/rungrawee-popcorn/hospital-middleware-system/internal/hospital"

	repository "github.com/rungrawee-popcorn/hospital-middleware-system/internal/repositories"
)

type PatientService struct {
	Repository     *repository.PatientRepository
	HospitalClient *hospital.HospitalClient
}

func NewPatientService(
	repository *repository.PatientRepository,
	hospitalClient *hospital.HospitalClient,
) *PatientService {

	return &PatientService{
		Repository:     repository,
		HospitalClient: hospitalClient,
	}
}

// SearchPatient searches patient information.
// Flow:
// 1. Search local database first.
// 2. If not found, call Hospital A API when national_id/passport_id exists.
// 3. Save external patient data into local database.
// 4. Return result.
func (s *PatientService) SearchPatient(
	hospitalID uint,
	filters map[string]interface{},
) ([]model.Patient, error) {

	if hospitalID == 0 {

		return nil, errors.New(
			"invalid hospital access",
		)
	}

	// Search local database first.
	patients, err := s.Repository.SearchPatients(
		hospitalID,
		filters,
	)

	if err != nil {

		return nil, err
	}

	// Found in local database.
	if len(patients) > 0 {

		return patients, nil
	}

	// Not found.
	// Hospital A API only supports national_id or passport_id.
	var identifier string

	if value, ok := filters["national_id"].(string); ok && value != "" {

		identifier = value

	}

	if identifier == "" {

		if value, ok := filters["passport_id"].(string); ok && value != "" {

			identifier = value

		}
	}

	// Cannot call Hospital A without identifier.
	if identifier == "" {

		return []model.Patient{}, nil
	}

	// Search from Hospital A.
	patient, err := s.SearchFromHospitalA(
		hospitalID,
		identifier,
	)

	if err != nil {

		return nil, err
	}

	return []model.Patient{
		*patient,
	}, nil
}

// SearchFromHospitalA searches patient from external hospital API.
func (s *PatientService) SearchFromHospitalA(
	hospitalID uint,
	identifier string,
) (*model.Patient, error) {

	if identifier == "" {

		return nil, errors.New(
			"patient identifier is required",
		)
	}

	// Call Hospital A API.
	response, err := s.HospitalClient.SearchPatient(
		identifier,
	)

	if err != nil {

		return nil, err
	}

	// Map external response to internal model.
	patient := model.Patient{

		HospitalID: hospitalID,

		PatientHN: response.PatientHN,

		NationalID: response.NationalID,

		PassportID: response.PassportID,

		FirstNameTH: response.FirstNameTH,

		MiddleNameTH: response.MiddleNameTH,

		LastNameTH: response.LastNameTH,

		FirstNameEN: response.FirstNameEN,

		MiddleNameEN: response.MiddleNameEN,

		LastNameEN: response.LastNameEN,

		DateOfBirth: response.DateOfBirth,

		PhoneNumber: response.PhoneNumber,

		Email: response.Email,

		Gender: response.Gender,
	}

	// Save patient data.
	err = s.Repository.CreatePatient(
		&patient,
	)

	if err != nil {

		return nil, err
	}

	return &patient, nil
}
