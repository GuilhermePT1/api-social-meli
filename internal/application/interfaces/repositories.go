package interfaces

import "github.com/GuilhermePT1/api-social-meli/internal/domain/models"

// USER
type UserRepository interface {
	Create(user *models.User) error
	FindById(id uint) (*models.User, error)
	FindAll() ([]models.User, error)
}

// FOLLOW
type FollowRepository interface {
	Create(follow *models.Follow) error
	Delete(userID, followerID uint) error
	CountFollowers(userID uint) (int64, error)
	FindFollowers(userID uint) ([]models.User, error)
	FindFollowed(userID uint) ([]models.User, error)
}

// PRODUCT

type ProductRepository interface {
	Create(product *models.Product) error
	FindById(id uint) (*models.Product, error)
	FindAll() ([]models.Product, error)
}

// POST

type PostRepository interface {
	Create(post *models.Post) error
	FindByUserLastTwoWeeks(userID uint) ([]models.Post, error)
	FindByUser(userID uint) ([]models.Post, error)
	FindPromoPosts() ([]models.Post, error)
	CountPromoProducts() (int64, error)
}
