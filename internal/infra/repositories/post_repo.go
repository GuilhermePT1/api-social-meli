package repositories

import (
	"time"

	"github.com/GuilhermePT1/api-social-meli/internal/domain/models"
	"gorm.io/gorm"
)

type PostRepository struct {
	DB *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{DB: db}
}

func (r *PostRepository) Create(post *models.Post) error {
	return r.DB.Create(post).Error
}

func (r *PostRepository) FindByUserLastTwoWeeks(userID uint) ([]models.Post, error) {
	var posts []models.Post
	twoWeeksAgo := time.Now().AddDate(0, 0, -14)

	err := r.DB.
		Where("user_id = ? AND created_at > ?", userID, twoWeeksAgo).
		Order("created_at DESC").
		Find(&posts).Error

	return posts, err
}

func (r *PostRepository) FindByUser(userID uint) ([]models.Post, error) {
	var posts []models.Post
	err := r.DB.Where("user_id = ?", userID).Find(&posts).Error
	return posts, err
}

func (r *PostRepository) FindPromoPosts() ([]models.Post, error) {
	var posts []models.Post

	err := r.DB.
		Where("has_promotion = ?", true).
		Find(&posts).Error

	return posts, err
}

func (r *PostRepository) CountPromoProducts() (int64, error) {
	var count int64

	err := r.DB.
		Model(&models.Post{}).
		Where("has_promotion = ?", true).
		Count(&count).Error

	return count, err
}
