package hospital

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type HospitalClient struct {
	BaseURL string
	Client  *http.Client
}

func NewHospitalClient(baseURL string) *HospitalClient {
	return &HospitalClient{
		BaseURL: baseURL,
		Client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

type HospitalPatientResponse struct {
	FirstNameTH  string `json:"first_name_th"`
	MiddleNameTH string `json:"middle_name_th"`
	LastNameTH   string `json:"last_name_th"`

	FirstNameEN  string `json:"first_name_en"`
	MiddleNameEN string `json:"middle_name_en"`
	LastNameEN   string `json:"last_name_en"`

	DateOfBirth time.Time `json:"date_of_birth"`

	PatientHN string `json:"patient_hn"`

	NationalID string `json:"national_id"`
	PassportID string `json:"passport_id"`

	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`

	Gender string `json:"gender"`
}

// SearchPatient calls Hospital A API.
func (h *HospitalClient) SearchPatient(
	identifier string,
) (*HospitalPatientResponse, error) {

	if identifier == "" {
		return nil, errors.New("patient identifier is required")
	}

	url := fmt.Sprintf(
		"%s/patient/search/%s",
		h.BaseURL,
		identifier,
	)

	request, err := http.NewRequest(
		http.MethodGet,
		url,
		nil,
	)

	if err != nil {
		return nil, err
	}

	response, err := h.Client.Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, errors.New("hospital api returned error")
	}

	var patient HospitalPatientResponse

	err = json.NewDecoder(
		response.Body,
	).Decode(&patient)

	if err != nil {
		return nil, err
	}

	return &patient, nil
}
