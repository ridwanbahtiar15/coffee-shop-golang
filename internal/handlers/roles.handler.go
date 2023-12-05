package handlers

import (
	"coffee-shop-golang/internal/helpers"
	"coffee-shop-golang/internal/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerRoles struct {
	*repositories.RolesRepository
}

func InitializeHandlerRoles(r *repositories.RolesRepository) *HandlerRoles {
	return &HandlerRoles{r}
}

func (h *HandlerRoles) GetAllRoles(ctx *gin.Context) {
	result, err := h.RepsitoryGetAllRoles()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if len(result) == 0 {
		ctx.JSON(http.StatusNotFound, helpers.GetResponse("role not found", nil))
		return
	}
	
	ctx.JSON(http.StatusOK, helpers.GetResponse("get all role success", result))
}