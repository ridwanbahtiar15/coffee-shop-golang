package handlers

import (
	"coffee-shop-golang/internal/helpers"
	"coffee-shop-golang/internal/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerCategories struct {
	repositories.ICategoriesRepository
}

func InitializeHandlerCategories(r repositories.ICategoriesRepository) *HandlerCategories {
	return &HandlerCategories{r}
}

func (h *HandlerCategories) GetAllCategories(ctx *gin.Context) {
	result, err := h.RepositoryGetAllCategories()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if len(result) == 0 {
		ctx.JSON(http.StatusNotFound, helpers.GetResponse("category not found", nil, nil))
		return
	}
	ctx.JSON(http.StatusOK, helpers.GetResponse("get all category success", result, nil))
}
