package handlers

import (
	"coffee-shop-golang/internal/helpers"
	"coffee-shop-golang/internal/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerCategories struct {
	*repositories.CategoriesRepository
}

func InitializeHandlerCategories(r *repositories.CategoriesRepository) *HandlerCategories {
	return &HandlerCategories{r}
}

func (h *HandlerCategories) GetAllCategories(ctx *gin.Context) {
	result, err := h.RepsitoryGetAllCategories()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if len(result) == 0 {
		ctx.JSON(http.StatusNotFound, helpers.GetResponse("category not found", nil))
		return
	}
	ctx.JSON(http.StatusOK, helpers.GetResponse("get all category success", result))
}
