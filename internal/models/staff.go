package model

import "time"

type Staff struct {

	ID uint `gorm:"primaryKey"`

	Username string `gorm:"uniqueIndex;not null"`

	Password string `gorm:"not null"`

	HospitalID uint `gorm:"not null;index"`

	Hospital Hospital `gorm:"foreignKey:HospitalID"`

	CreatedAt time.Time

	UpdatedAt time.Time
}