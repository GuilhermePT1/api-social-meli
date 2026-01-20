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

type UserController struct {
	Service *services.UserService
}

// @Summary Cria um novo usuário
// @Description Cria um novo usuário com os dados fornecidos
// @Tags User
// @Accept json
// @Produce json
// @Param user body dto.UserRequestDTO true "Informações do usuário"
// @Success 201 {object} dto.UserResponseDTO
// @Router /api/users [post]
func NewUserController(service *services.UserService) *UserController {
	return &UserController{Service: service}
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var req dto.UserRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := c.Service.Create(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, dto.UserResponseDTO{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	})
}

// @Summary Listar um usuário por ID
// @Description Lista um usuário específico por ID
// @Tags User
// @Accept json
// @Produce json
// @Param user_id path string true "ID do usuário"
// @Success 200 {object} dto.UserResponseDTO
// @Router /api/users/{user_id} [get]
func (c *UserController) GetUserById(ctx *gin.Context) {
	idStr := ctx.Param("user_id")
	if idStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "É necessário preencher o id do usuário"})
		return
	}

	id64, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID do usuário inválido"})
		return
	}

	id := uint(id64)
	u, err := c.Service.GetById(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar usuário"})
		return
	}

	ctx.JSON(http.StatusOK, u)
}

// @Summary Listar todos os usuários
// @Description Lista todos os usuários
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} []dto.UserResponseDTO
// @Router /api/users [get]
func (c *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := c.Service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}
