package controllers

import (
	"net/http"
	"strconv"

	"github.com/GuilhermePT1/api-social-meli/internal/application/services"
	"github.com/GuilhermePT1/api-social-meli/internal/domain/dto"
	"github.com/gin-gonic/gin"
)

type FollowController struct {
	Service *services.FollowService
}

// @Summary Seguir um usuário
// @Description Segue um usuário específico
// @Tags Follow
// @Accept json
// @Produce json
// @Param follow body dto.FollowRequestDTO true "Informações do usuário a ser seguido"
// @Success 200 {object} dto.FollowResponseDTO
// @Router /api/users/follow [post]
func NewFollowController(service *services.FollowService) *FollowController {
	return &FollowController{Service: service}
}

func (c *FollowController) Follow(ctx *gin.Context) {
	var req dto.FollowRequestDTO

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.Service.Follow(req.UserID, req.FollowerID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Seguindo com sucesso"})
}

// @Summary Deixar de seguir um usuário
// @Description Deixa de seguir um usuário específico
// @Tags Follow
// @Accept json
// @Produce json
// @Param follow body dto.FollowRequestDTO true "Informações do usuário a ser deixado de seguir"
// @Success 200 {object} dto.FollowResponseDTO
// @Router /api/users/unfollow [post]
func (c *FollowController) Unfollow(ctx *gin.Context) {
	var req dto.FollowRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.Service.Unfollow(req.UserID, req.FollowerID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Deixando de seguir com sucesso"})
}

// @Summary Contar seguidores de um usuário
// @Description Conta o número de seguidores de um usuário específico
// @Tags Follow
// @Accept json
// @Produce json
// @Param user_id path string true "ID do usuário"
// @Success 200 {object} []models.User
// @Router /api/users/{user_id}/followers/count [get]
func (c *FollowController) CountFollowers(ctx *gin.Context) {
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

	count, err := c.Service.CountFollowers(uint(userID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Quantidade de seguidores", "count": count})
}

// @Summary Listar seguidores de um usuário
// @Description Lista os seguidores de um usuário específico
// @Tags Follow
// @Accept json
// @Produce json
// @Param user_id path string true "ID do usuário"
// @Success 200 {object} []models.User
// @Router /api/users/{user_id}/followers/list [get]
func (c *FollowController) GetFollowers(ctx *gin.Context) {
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

	followers, err := c.Service.GetFollowers(uint(userID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, followers)
}

// @Summary Listar usuários seguidos por um usuário
// @Description Lista os usuários seguidos por um usuário específico
// @Tags Follow
// @Accept json
// @Produce json
// @Param user_id path string true "ID do usuário"
// @Success 200 {object} []models.User
// @Router /api/users/{user_id}/followed/list [get]
func (c *FollowController) GetFollowed(ctx *gin.Context) {
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

	followed, err := c.Service.GetFollowed(uint(userID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, followed)
}
