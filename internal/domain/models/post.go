package models

import "time"

type Post struct {
	ID           uint `gorm:"primaryKey"`
	UserID       uint `gorm:"index"`
	ProductID    uint `gorm:"index"`
	Category     int
	Price        float64
	HasPromotion bool
	Discount     float64
	CreatedAt    time.Time `gorm:"index"`

	Product Product `gorm:"foreignKey:ProductID"`
}
