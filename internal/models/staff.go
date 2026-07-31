package model

import "time"

type Staff struct {
	ID uint `gorm:"primaryKey"`

	// Username must be unique within the same hospital.
	Username string `gorm:"not null;uniqueIndex:idx_staff_username_hospital"`

	// Password is stored as a bcrypt hash.
	Password string `gorm:"not null"`

	// Composite unique index with Username.
	HospitalID uint `gorm:"not null;index;uniqueIndex:idx_staff_username_hospital"`

	Hospital Hospital `gorm:"foreignKey:HospitalID"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
