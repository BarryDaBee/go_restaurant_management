package models

import "time"

type Order struct {
	ID        uint    `gorm:"primaryKey"`
	UserID    uint    `gorm:"not null"`
	Total     float64 `gorm:"not null"`
	Status    string  `gorm:"not null;default:'pending'"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
