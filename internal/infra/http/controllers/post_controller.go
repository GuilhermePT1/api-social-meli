package controllers

import (
	"net/http"
	"strconv"

	"github.com/GuilhermePT1/api-social-meli/internal/application/services"
	"github.com/GuilhermePT1/api-social-meli/internal/domain/dto"
	"github.com/gin-gonic/gin"
)

type PostController struct {
	Service *services.PostService
}

func NewPostController(service *services.PostService) *PostController {
	return &PostController{Service: service}
}

func (c *PostController) CreatePost(ctx *gin.Context) {
	var req dto.PostRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post, err := c.Service.Create(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, dto.PostResponseDTO{
		ID:        post.ID,
		UserID:    post.UserID,
		ProductID: post.ProductID,
		Price:     post.Price,
		Promotion: post.HasPromotion,
		Discount:  post.Discount,
	})
}

func (c *PostController) FindByUser(ctx *gin.Context) {
	userIDStr := ctx.Param("user_id")
	if userIDStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "É necessário preencher o id do úsuario"})
		return
	}

	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID do usuário inválido"})
		return
	}

	posts, err := c.Service.FindByUser(uint(userID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var response []dto.PostResponseDTO
	for _, post := range posts {
		response = append(response, dto.PostResponseDTO{
			ID:        post.ID,
			UserID:    post.UserID,
			ProductID: post.ProductID,
			Price:     post.Price,
			Promotion: post.HasPromotion,
			Discount:  post.Discount,
		})
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *PostController) FindPromoPosts(ctx *gin.Context) {
	posts, err := c.Service.FindPromoPosts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var response []dto.PostResponseDTO
	for _, post := range posts {
		response = append(response, dto.PostResponseDTO{
			ID:        post.ID,
			UserID:    post.UserID,
			ProductID: post.ProductID,
			Price:     post.Price,
			Promotion: post.HasPromotion,
			Discount:  post.Discount,
		})
	}
	ctx.JSON(http.StatusOK, response)
}

func (c *PostController) CountPromoProducts(ctx *gin.Context) {
	count, err := c.Service.CountPromoProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Quantidade de produtos com promoção", "count": count})
}
