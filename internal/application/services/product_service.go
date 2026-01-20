package services

import (
	"github.com/GuilhermePT1/api-social-meli/internal/application/interfaces"
	"github.com/GuilhermePT1/api-social-meli/internal/domain/dto"
	"github.com/GuilhermePT1/api-social-meli/internal/domain/models"
)

type ProductService struct {
	Repo interfaces.ProductRepository
}

func NewProductService(repo interfaces.ProductRepository) *ProductService {
	return &ProductService{Repo: repo}
}

func (s *ProductService) Create(req dto.ProductRequestDTO) (*models.Product, error) {
	product := &models.Product{
		Name:  req.Name,
		Type:  req.Type,
		Brand: req.Brand,
		Color: req.Color,
		Notes: req.Notes,
	}
	err := s.Repo.Create(product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *ProductService) GetById(id uint) (*models.Product, error) {
	return s.Repo.FindById(id)
}

func (s *ProductService) GetAll() ([]models.Product, error) {
	return s.Repo.FindAll()
}
