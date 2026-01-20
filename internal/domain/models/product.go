package models

import "time"

type Product struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:255;not null"`
	Type      string
	Brand     string
	Color     string
	Notes     string
	CreatedAt time.Time `gorm:"index"`
}
