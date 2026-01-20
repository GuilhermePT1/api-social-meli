package services

import (
	"github.com/GuilhermePT1/api-social-meli/internal/application/interfaces"
	"github.com/GuilhermePT1/api-social-meli/internal/domain/models"
)

type FollowService struct {
	Repo interfaces.FollowRepository
}

func NewFollowService(repo interfaces.FollowRepository) *FollowService {
	return &FollowService{Repo: repo}
}

func (s *FollowService) Follow(userID, followerID uint) error {
	follow := &models.Follow{
		UserID:     userID,
		FollowerID: followerID,
	}
	return s.Repo.Create(follow)
}

func (s *FollowService) Unfollow(userID, followerID uint) error {
	return s.Repo.Delete(userID, followerID)
}

func (s *FollowService) CountFollowers(userID uint) (int64, error) {
	return s.Repo.CountFollowers(userID)
}

func (s *FollowService) GetFollowers(userID uint) ([]models.User, error) {
	return s.Repo.FindFollowers(userID)
}

func (s *FollowService) GetFollowed(userID uint) ([]models.User, error) {
	return s.Repo.FindFollowed(userID)
}
