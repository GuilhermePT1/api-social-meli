package dto

type ProductRequestDTO struct {
	Name  string `json:"name" binding:"required"`
	Type  string `json:"type" binding:"required"`
	Brand string `json:"brand" binding:"required"`
	Color string `json:"color" binding:"required"`
	Notes string `json:"notes" binding:"required"`
}

type ProductResponseDTO struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Type  string `json:"type"`
	Brand string `json:"brand"`
	Color string `json:"color"`
	Notes string `json:"notes"`
}
