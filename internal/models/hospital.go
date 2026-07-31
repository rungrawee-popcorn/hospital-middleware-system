package model

import "time"

type Hospital struct {
	ID uint `gorm:"primaryKey"`

	Name string `gorm:"not null;uniqueIndex"`

	CreatedAt time.Time

	UpdatedAt time.Time
}
