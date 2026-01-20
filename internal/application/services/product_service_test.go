package services_test

import (
	"testing"

	"github.com/GuilhermePT1/api-social-meli/internal/application/services"
	"github.com/GuilhermePT1/api-social-meli/internal/domain/dto"
	"github.com/GuilhermePT1/api-social-meli/internal/domain/models"
	"github.com/stretchr/testify/assert"
)

type productRepoMock struct {
	createCalled   bool
	findAllCalled  bool
	findByIdCalled bool

	createInput   models.Product
	findByIdInput uint

	productsResult []models.Product
	product        *models.Product
	err            error
}

func (m *productRepoMock) Create(product *models.Product) error {
	m.createCalled = true
	m.createInput = *product
	return m.err
}

func (m *productRepoMock) FindById(id uint) (*models.Product, error) {
	m.findByIdCalled = true
	m.findByIdInput = id
	return m.product, m.err
}

func (m *productRepoMock) FindAll() ([]models.Product, error) {
	m.findAllCalled = true
	return m.productsResult, m.err
}

func TestProductService_Create_Success(t *testing.T) {
	repo := &productRepoMock{}
	service := services.NewProductService(repo)

	req := dto.ProductRequestDTO{
		Name:  "MacBook Pro",
		Type:  "Laptop",
		Brand: "Apple",
		Color: "Silver",
		Notes: "Apple's latest laptop",
	}

	created, err := service.Create(req)

	assert.NoError(t, err)
	assert.NotNil(t, created)
	assert.True(t, repo.createCalled)
}

func TestProductService_Create_Error(t *testing.T) {
	repo := &productRepoMock{err: assert.AnError}
	service := services.NewProductService(repo)

	req := dto.ProductRequestDTO{
		Name:  "",
		Type:  "Laptop",
		Brand: "Apple",
		Color: "Silver",
		Notes: "Apple's latest laptop",
	}

	_, err := service.Create(req)

	assert.Error(t, err)
	assert.Nil(t, nil)
	assert.True(t, repo.createCalled)
}

func TestProductService_FindAll_Success(t *testing.T) {
	repo := &productRepoMock{
		productsResult: []models.Product{
			{ID: 1, Name: "MacBook Pro", Type: "Laptop", Brand: "Apple", Color: "Silver", Notes: "Apple's latest laptop"},
			{ID: 2, Name: "iPhone 13", Type: "Phone", Brand: "Apple", Color: "Black", Notes: "Apple's latest phone"},
		},
	}

	service := services.NewProductService(repo)

	products, err := service.GetAll()

	assert.NoError(t, err)
	assert.NotNil(t, products)
	assert.True(t, repo.findAllCalled)
}

func TestProductService_FindAll_Error(t *testing.T) {
	repo := &productRepoMock{err: assert.AnError}
	service := services.NewProductService(repo)

	products, err := service.GetAll()

	assert.Error(t, err)
	assert.Nil(t, products)
}

func TestProductService_FindById_Success(t *testing.T) {
	repo := &productRepoMock{
		product: &models.Product{ID: 1, Name: "MacBook Pro", Type: "Laptop", Brand: "Apple", Color: "Silver", Notes: "Apple's latest laptop"},
	}
	service := services.NewProductService(repo)

	product, err := service.GetById(1)

	assert.NoError(t, err)
	assert.NotNil(t, product)
	assert.True(t, repo.findByIdCalled)
}

func TestProductService_FindById_Error(t *testing.T) {
	repo := &productRepoMock{err: assert.AnError}
	service := services.NewProductService(repo)

	product, err := service.GetById(1)

	assert.Error(t, err)
	assert.Nil(t, product)
}
