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

// @Summary Cria um novo post
// @Description Cria um novo post com os dados fornecidos
// @Tags Post
// @Accept json
// @Produce json
// @Param post body dto.PostRequestDTO true "Informações do post"
// @Success 201 {object} dto.PostResponseDTO
// @Router /api/posts [post]
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

// @Summary Listar posts de um usuário
// @Description Lista os posts de um usuário específico
// @Tags Post
// @Accept json
// @Produce json
// @Param user_id path string true "ID do usuário"
// @Success 200 {object} []dto.PostResponseDTO
// @Router /api/users/{user_id}/posts [get]
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

// @Summary Listar posts com promoção
// @Description Lista os posts com promoção
// @Tags Post
// @Accept json
// @Produce json
// @Success 200 {object} []dto.PostResponseDTO
// @Router /api/posts/promo [get]
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

// @Summary Contar produtos com promoção
// @Description Conta o número de produtos com promoção
// @Tags Post
// @Accept json
// @Produce json
// @Success 200 {object} []dto.PostResponseDTO
// @Router /api/posts/promo/count [get]
func (c *PostController) CountPromoProducts(ctx *gin.Context) {
	count, err := c.Service.CountPromoProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Quantidade de produtos com promoção", "count": count})
}
