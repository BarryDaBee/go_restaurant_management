package models

import "time"

type Food struct {
	ID        uint    `gorm:"primaryKey"`
	Name      string  `gorm:"not null"`
	Price     float64 `gorm:"not null"`
	Category  string  `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
