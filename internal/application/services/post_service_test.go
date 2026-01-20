package services_test

import (
	"testing"

	"github.com/GuilhermePT1/api-social-meli/internal/application/services"
	"github.com/GuilhermePT1/api-social-meli/internal/domain/dto"
	"github.com/GuilhermePT1/api-social-meli/internal/domain/models"
	"github.com/stretchr/testify/assert"
)

type postRepoMock struct {
	createCalled                 bool
	findByUserLastTwoWeeksCalled bool
	findByUserCalled             bool
	findPromoPostCalled          bool
	CountPromoProductsCalled     bool

	createInput models.Post
	userID      uint

	posts []models.Post
	count int64
	err   error
}

func (m *postRepoMock) Create(post *models.Post) error {
	m.createCalled = true
	m.createInput = *post
	return m.err
}

func (m *postRepoMock) FindByUserLastTwoWeeks(userID uint) ([]models.Post, error) {
	m.findByUserLastTwoWeeksCalled = true
	m.userID = userID
	return m.posts, m.err
}

func (m *postRepoMock) FindByUser(userID uint) ([]models.Post, error) {
	m.findByUserCalled = true
	m.userID = userID
	return m.posts, m.err
}

func (m *postRepoMock) FindPromoPosts() ([]models.Post, error) {
	m.findPromoPostCalled = true
	return m.posts, m.err
}

func (m *postRepoMock) CountPromoProducts() (int64, error) {
	m.CountPromoProductsCalled = true
	return m.count, m.err
}

func TestPostService_Create_Success(t *testing.T) {
	repo := &postRepoMock{}
	service := services.NewPostService(repo)

	created, err := service.Create(dto.PostRequestDTO{
		UserID:    1,
		ProductID: 1,
		Price:     100,
		Promotion: true,
	})

	assert.NoError(t, err)
	assert.NotNil(t, created)
	assert.True(t, repo.createCalled)
}

func TestPostService_Create_Error(t *testing.T) {
	repo := &postRepoMock{err: assert.AnError}
	service := services.NewPostService(repo)

	created, err := service.Create(dto.PostRequestDTO{
		UserID:    1,
		ProductID: 1,
		Price:     100,
		Promotion: true,
	})

	assert.Error(t, err)
	assert.Nil(t, created)
}

func TestPostService_FindByUser_Success(t *testing.T) {
	repo := &postRepoMock{
		posts: []models.Post{
			{ID: 1, UserID: 1, ProductID: 1, Price: 100, HasPromotion: true, Discount: 10},
		},
	}
	service := services.NewPostService(repo)

	posts, err := service.FindByUser(1)

	assert.NoError(t, err)
	assert.NotNil(t, posts)
	assert.True(t, repo.findByUserCalled)
}

func TestPostService_FindByUser_Error(t *testing.T) {
	repo := &postRepoMock{err: assert.AnError}
	service := services.NewPostService(repo)

	posts, err := service.FindByUser(1)

	assert.Error(t, err)
	assert.Nil(t, posts)
}

func TestPostService_FindPromoPosts_Success(t *testing.T) {
	repo := &postRepoMock{
		posts: []models.Post{
			{ID: 1, UserID: 1, ProductID: 1, Price: 100, HasPromotion: true, Discount: 10},
		},
	}
	service := services.NewPostService(repo)

	posts, err := service.FindPromoPosts()

	assert.NoError(t, err)
	assert.NotNil(t, posts)
	assert.True(t, repo.findPromoPostCalled)
}

func TestPostService_FindPromoPosts_Error(t *testing.T) {
	repo := &postRepoMock{err: assert.AnError}
	service := services.NewPostService(repo)

	posts, err := service.FindPromoPosts()

	assert.Error(t, err)
	assert.Nil(t, posts)
}
