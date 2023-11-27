package handlers

import (
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
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "category not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
		"message": "get all category success",
	})
}
