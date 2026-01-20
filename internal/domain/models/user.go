package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"size:255;not null"`
	Email     string    `gorm:"size:255;not null;unique"`
	Role      string    `gorm:"size:255;not null"`
	CreatedAt time.Time `gorm:"index"`
}
