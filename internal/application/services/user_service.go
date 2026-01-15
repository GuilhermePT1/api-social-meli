package services

import (
	"github.com/GuilhermePT1/api-social-meli/internal/application/interfaces"
	"github.com/GuilhermePT1/api-social-meli/internal/domain/dto"
	"github.com/GuilhermePT1/api-social-meli/internal/domain/models"
)

type UserService struct {
	Repo interfaces.UserRepository
}

func NewUserService(repo interfaces.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) Create(req dto.UserRequestDTO) (*models.User, error) {
	user := &models.User{
		Name:  req.Name,
		Email: req.Email,
		Role:  req.Role,
	}
	err := s.Repo.Create(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetById(id uint) (*models.User, error) {
	return s.Repo.FindById(id)
}

func (s *UserService) GetAll() ([]models.User, error) {
	return s.Repo.FindAll()
}
