package handlers

import (
	"coffee-shop-golang/internal/models"
	"coffee-shop-golang/internal/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerCategories struct {
	*repositories.CategoriesRepository
}

func InitializeHandler(r *repositories.CategoriesRepository) *HandlerCategories {
	return &HandlerCategories{r}
}

func (h *HandlerCategories) GetAllCategories(ctx *gin.Context) {
	result, err := h.RepsitoryGetAllCategories()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
		"message": "get all category success",
	})
}

func (h *HandlerCategories) CreateCategories(ctx *gin.Context) {
	var body models.CategoriesModel
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	result, err := h.RepsitoryCreateCategories(&body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
		"message": "create category success",
	})
}

func (h *HandlerCategories) UpdateCategories(ctx *gin.Context) {
	id := ctx.Param("id")

	var body models.CategoriesModel
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	result, err := h.RepsitoryUpdateCategories(&body, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
		"message": "update category success",
	})
}

func (h *HandlerCategories) DeleteCategories(ctx *gin.Context) {
	id := ctx.Param("id")
	result, err := h.RepositoryDeleteCategories(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
		"message": "delete category success",
	})
}