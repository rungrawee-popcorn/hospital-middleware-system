package model

import "time"

type Patient struct {

	ID uint `gorm:"primaryKey"`

	HospitalID uint `gorm:"not null;index"`

	Hospital Hospital `gorm:"foreignKey:HospitalID"`


	PatientHN string `gorm:"uniqueIndex"`

	NationalID string `gorm:"index"`

	PassportID string `gorm:"index"`


	FirstNameTH string

	MiddleNameTH string

	LastNameTH string


	FirstNameEN string

	MiddleNameEN string

	LastNameEN string


	DateOfBirth time.Time


	PhoneNumber string `gorm:"index"`

	Email string `gorm:"index"`


	Gender string `gorm:"type:varchar(1)"`


	CreatedAt time.Time

	UpdatedAt time.Time
}