package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/GuilhermePT1/api-social-meli/internal/application/services"
	"github.com/GuilhermePT1/api-social-meli/internal/domain/dto"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductController struct {
	Service *services.ProductService
}

// @Summary Cria um novo produto
// @Description Cria um novo produto com os dados fornecidos
// @Tags Product
// @Accept json
// @Produce json
// @Param product body dto.ProductRequestDTO true "Informações do produto"
// @Success 201 {object} dto.ProductResponseDTO
// @Router /api/products [post]
func NewProductController(service *services.ProductService) *ProductController {
	return &ProductController{Service: service}
}

func (c *ProductController) CreateProduct(ctx *gin.Context) {
	var req dto.ProductRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := c.Service.Create(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, dto.ProductResponseDTO{
		ID:    product.ID,
		Name:  product.Name,
		Type:  product.Type,
		Brand: product.Brand,
		Color: product.Color,
		Notes: product.Notes,
	})
}

// @Summary Listar um produto por ID
// @Description Lista um produto específico por ID
// @Tags Product
// @Accept json
// @Produce json
// @Param product_id path string true "ID do produto"
// @Success 200 {object} dto.ProductResponseDTO
// @Router /api/products/{product_id} [get]
func (c *ProductController) GetProductById(ctx *gin.Context) {
	idStr := ctx.Param("product_id")
	if idStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "É necessário preencher o id do produto"})
		return
	}

	id64, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID do produto inválido"})
		return
	}

	id := uint(id64)
	p, err := c.Service.GetById(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Produto não encontrado"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar produto"})
		return
	}

	ctx.JSON(http.StatusOK, p)
}

// @Summary Listar todos os produtos
// @Description Lista todos os produtos
// @Tags Product
// @Accept json
// @Produce json
// @Success 200 {object} []dto.ProductResponseDTO
// @Router /api/products [get]
func (c *ProductController) GetAllProducts(ctx *gin.Context) {
	products, err := c.Service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, products)
}
