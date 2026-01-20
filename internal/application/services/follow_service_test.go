package services_test

import (
	"testing"

	"github.com/GuilhermePT1/api-social-meli/internal/application/interfaces"
	"github.com/GuilhermePT1/api-social-meli/internal/application/services"
	"github.com/GuilhermePT1/api-social-meli/internal/domain/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type followRepoMock struct {
	mock.Mock
	CreateCalled bool
	DeleteCalled bool

	createInput    models.Follow
	deleteUser     uint
	deleteFollower uint

	err error

	// extras exigidos pela interface
	CountFollowersCalled bool
	countFollowersUserID uint
	countFollowersResult int64
	countFollowersErr    error

	FindFollowersCalled bool
	findFollowersUserID uint
	findFollowersResult []models.User
	findFollowersErr    error

	FindFollowedCalled bool
	findFollowedUserID uint
	findFollowedResult []models.User
	findFollowedErr    error
}

var _ interfaces.FollowRepository = (*followRepoMock)(nil)

func (m *followRepoMock) Create(follow *models.Follow) error {
	m.CreateCalled = true
	m.createInput = *follow
	return m.err
}

func (m *followRepoMock) Delete(userID, followerID uint) error {
	m.DeleteCalled = true
	m.deleteUser = userID
	m.deleteFollower = followerID
	return m.err
}

func (m *followRepoMock) CountFollowers(userID uint) (int64, error) {
	m.CountFollowersCalled = true
	m.countFollowersUserID = userID
	return m.countFollowersResult, m.countFollowersErr
}

func (m *followRepoMock) FindFollowers(userID uint) ([]models.User, error) {
	m.FindFollowersCalled = true
	m.findFollowersUserID = userID
	return m.findFollowersResult, m.findFollowersErr
}

func (m *followRepoMock) FindFollowed(userID uint) ([]models.User, error) {
	m.FindFollowedCalled = true
	m.findFollowedUserID = userID
	return m.findFollowedResult, m.findFollowedErr
}

// testes

func TestFollowService_Follow_Success(t *testing.T) {
	repo := &followRepoMock{}
	service := services.NewFollowService(repo)

	err := service.Follow(1, 2)

	assert.NoError(t, err)
	assert.True(t, repo.CreateCalled)
	assert.Equal(t, uint(1), repo.createInput.UserID)
	assert.Equal(t, uint(2), repo.createInput.FollowerID)
}

func TestFollowService_Unfollow_Success(t *testing.T) {
	repo := &followRepoMock{}
	service := services.NewFollowService(repo)

	err := service.Unfollow(1, 2)

	assert.NoError(t, err)
	assert.True(t, repo.DeleteCalled)
	assert.Equal(t, uint(1), repo.deleteUser)
	assert.Equal(t, uint(2), repo.deleteFollower)
}

func TestFollowService_Follow_Error(t *testing.T) {
	repo := &followRepoMock{err: assert.AnError}
	service := services.NewFollowService(repo)

	err := service.Follow(1, 2)
	assert.Error(t, err)
}

func TestFollowService_Unfollow_Error(t *testing.T) {
	repo := &followRepoMock{err: assert.AnError}
	service := services.NewFollowService(repo)

	err := service.Unfollow(1, 2)
	assert.Error(t, err)
}

func TestFollowService_CountFollowers_Success(t *testing.T) {
	repo := &followRepoMock{}
	service := services.NewFollowService(repo)

	count, err := service.CountFollowers(1)
	assert.NoError(t, err)
	assert.True(t, repo.CountFollowersCalled)
	assert.Equal(t, uint(1), repo.countFollowersUserID)
	assert.Equal(t, int64(0), count)
}
