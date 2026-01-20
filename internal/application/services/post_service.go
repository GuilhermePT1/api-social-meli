package services

import (
	"github.com/GuilhermePT1/api-social-meli/internal/application/interfaces"
	"github.com/GuilhermePT1/api-social-meli/internal/domain/dto"
	"github.com/GuilhermePT1/api-social-meli/internal/domain/models"
)

type PostService struct {
	Repo interfaces.PostRepository
}

func NewPostService(repo interfaces.PostRepository) *PostService {
	return &PostService{Repo: repo}
}

func (s *PostService) Create(req dto.PostRequestDTO) (*models.Post, error) {
	discount := req.Price - req.Discount
	if !req.Promotion {
		discount = 0
	}

	post := &models.Post{
		UserID:       req.UserID,
		ProductID:    req.ProductID,
		Price:        req.Price,
		HasPromotion: req.Promotion,
		Discount:     discount,
	}
	err := s.Repo.Create(post)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (s *PostService) FindByUserLastTwoWeeks(userID uint) ([]models.Post, error) {
	return s.Repo.FindByUserLastTwoWeeks(userID)
}

func (s *PostService) FindByUser(userID uint) ([]models.Post, error) {
	return s.Repo.FindByUser(userID)
}

func (s *PostService) FindPromoPosts() ([]models.Post, error) {
	return s.Repo.FindPromoPosts()
}

func (s *PostService) CountPromoProducts() (int64, error) {
	return s.Repo.CountPromoProducts()
}
