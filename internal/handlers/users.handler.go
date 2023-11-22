package handlers

import (
	"coffee-shop-golang/internal/models"
	"coffee-shop-golang/internal/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerUsers struct {
	*repositories.UsersRepository
}

func InitializeHandlerUsers(r *repositories.UsersRepository) *HandlerUsers {
	return &HandlerUsers{r}
}

func (h *HandlerUsers) GetAllUsers(ctx *gin.Context) {
	result, err := h.RepsitoryGetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
		"message": "get all user success",
	})
}

func (h *HandlerUsers) GetUsersById(ctx *gin.Context) {
	id := ctx.Param("id")
	result, err := h.RepsitoryUsersById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
		"message": "get user by id success",
	})
}

func (h *HandlerUsers) CreateUsers(ctx *gin.Context) {
	var body models.UsersModel
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	result, err := h.RepsitoryCreateUsers(&body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
		"message": "create user success",
	})
}

func (h *HandlerUsers) UpdateUsers(ctx *gin.Context) {
	id := ctx.Param("id")

	var body models.UsersModel
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	result, err := h.RepsitoryUpdateUsers(&body, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
		"message": "update user success",
	})
}

func (h *HandlerUsers) DeleteUsers(ctx *gin.Context) {
	id := ctx.Param("id")
	result, err := h.RepositoryDeleteUsers(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
		"message": "delete user success",
	})
}