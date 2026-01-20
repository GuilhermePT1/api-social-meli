package services_test

import (
	"testing"

	"github.com/GuilhermePT1/api-social-meli/internal/application/services"
	"github.com/GuilhermePT1/api-social-meli/internal/domain/dto"
	"github.com/GuilhermePT1/api-social-meli/internal/domain/models"
	"github.com/stretchr/testify/assert"
)

type userRepoMock struct {
	createCalled   bool
	findByIdCalled bool
	findAllCalled  bool

	createInput   models.User
	findByIdInput uint

	users []models.User
	user  *models.User
	err   error
}

func (m *userRepoMock) Create(user *models.User) error {
	m.createCalled = true
	m.createInput = *user
	return m.err
}

func (m *userRepoMock) FindById(id uint) (*models.User, error) {
	m.findByIdCalled = true
	m.findByIdInput = id
	return m.user, m.err
}

func (m *userRepoMock) FindAll() ([]models.User, error) {
	m.findAllCalled = true
	return m.users, m.err
}

func TestUserService_Create_Success(t *testing.T) {
	repo := &userRepoMock{}
	service := services.NewUserService(repo)

	req := dto.UserRequestDTO{
		Name:  "John Doe",
		Email: "john.doe@example.com",
		Role:  "user",
	}

	created, err := service.Create(req)

	assert.NoError(t, err)
	assert.NotNil(t, created)
	assert.True(t, repo.createCalled)
}

func TestUserService_Create_Error(t *testing.T) {
	repo := &userRepoMock{err: assert.AnError}
	service := services.NewUserService(repo)

	req := dto.UserRequestDTO{
		Name:  "",
		Email: "john.doe@example.com",
		Role:  "user",
	}

	_, err := service.Create(req)

	assert.Error(t, err)
	assert.Nil(t, nil)
}

func TestUserService_FindById_Success(t *testing.T) {
	repo := &userRepoMock{
		user: &models.User{ID: 1, Name: "John Doe", Email: "john.doe@example.com", Role: "user"},
	}
	service := services.NewUserService(repo)

	user, err := service.GetById(1)

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.True(t, repo.findByIdCalled)
}

func TestUserService_FindById_Error(t *testing.T) {
	repo := &userRepoMock{err: assert.AnError}
	service := services.NewUserService(repo)

	user, err := service.GetById(1)

	assert.Error(t, err)
	assert.Nil(t, user)
}

func TestUserService_FindAll_Success(t *testing.T) {
	repo := &userRepoMock{
		users: []models.User{
			{ID: 1, Name: "John Doe", Email: "john.doe@example.com", Role: "user"},
			{ID: 2, Name: "Jane Doe", Email: "jane.doe@example.com", Role: "user"},
		},
	}
	service := services.NewUserService(repo)

	users, err := service.GetAll()

	assert.NoError(t, err)
	assert.NotNil(t, users)
	assert.True(t, repo.findAllCalled)
}

func TestUserService_FindAll_Error(t *testing.T) {
	repo := &userRepoMock{err: assert.AnError}
	service := services.NewUserService(repo)

	users, err := service.GetAll()

	assert.Error(t, err)
	assert.Nil(t, users)
}
