package repositories

import (
	"github.com/GuilhermePT1/api-social-meli/internal/domain/models"
	"gorm.io/gorm"
)

type FollowRepository struct {
	DB *gorm.DB
}

func NewFollowRepository(db *gorm.DB) *FollowRepository {
	return &FollowRepository{DB: db}
}

func (r *FollowRepository) Create(follow *models.Follow) error {
	return r.DB.Create(follow).Error
}

func (r *FollowRepository) Delete(userID, followerID uint) error {
	return r.DB.Where("user_id = ? AND follower_id = ?", userID, followerID).Delete(&models.Follow{}).Error
}

func (r *FollowRepository) CountFollowers(userID uint) (int64, error) {
	var count int64

	err := r.DB.
		Model(&models.Follow{}).
		Where("user_id = ?", userID).
		Count(&count).Error

	return count, err
}

func (r *FollowRepository) FindFollowers(userID uint) ([]models.User, error) {
	var users []models.User

	err := r.DB.
		Table("users").
		Joins("JOIN follows ON follows.follower_id = users.id").
		Where("follows.user_id = ?", userID).
		Scan(&users).Error

	return users, err
}

func (r *FollowRepository) FindFollowed(userID uint) ([]models.User, error) {
	var users []models.User

	err := r.DB.
		Table("users").
		Joins("JOIN follows ON follows.user_id = users.id").
		Where("follows.follower_id = ?", userID).
		Scan(&users).Error

	return users, err
}
