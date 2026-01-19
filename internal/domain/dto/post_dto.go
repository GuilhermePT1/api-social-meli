package dto

import "time"

type PostRequestDTO struct {
	UserID    uint    `json:"user_id" binding:"required"`
	ProductID uint    `json:"product_id" binding:"required"`
	Price     float64 `json:"price" binding:"required"`
	Promotion bool    `json:"promotion" binding:"required"`
	Discount  float64 `json:"discount" binding:"required"`
}

type PostResponseDTO struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	ProductID uint      `json:"product_id"`
	Price     float64   `json:"price"`
	Promotion bool      `json:"promotion"`
	Discount  float64   `json:"discount"`
	CreatedAt time.Time `json:"created_at"`
}
