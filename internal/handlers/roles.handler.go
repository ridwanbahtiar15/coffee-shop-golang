package handlers

import (
	"coffee-shop-golang/internal/helpers"
	"coffee-shop-golang/internal/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerRoles struct {
	repositories.IRoleRepository
}

func InitializeHandlerRoles(r repositories.IRoleRepository) *HandlerRoles {
	return &HandlerRoles{r}
}

func (h *HandlerRoles) GetAllRoles(ctx *gin.Context) {
	result, err := h.RepositoryGetAllRoles()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if len(result) == 0 {
		ctx.JSON(http.StatusNotFound, helpers.GetResponse("role not found", nil, nil))
		return
	}
	
	ctx.JSON(http.StatusOK, helpers.GetResponse("get all role success", result, nil))
}