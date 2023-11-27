package handlers

import (
	"coffee-shop-golang/internal/models"
	"coffee-shop-golang/internal/repositories"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerPromos struct {
	*repositories.PromosRepository
}

func InitializeHandlerPromos(r *repositories.PromosRepository) *HandlerPromos {
	return &HandlerPromos{r}
}

func (h *HandlerPromos) GetAllPromos(ctx *gin.Context) {
	result, err := h.RepsitoryGetAllPromos()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
		"message": "get all promo success",
	})
}

func (h *HandlerPromos) CreateProomos(ctx *gin.Context) {
	var body models.PromosModel
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	result, err := h.RepsitoryCreatePromos(&body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
		"message": "create promo success",
	})
}

func (h *HandlerPromos) UpdatePromos(ctx *gin.Context) {
	id := ctx.Param("id")

	var body models.PromosModel
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	result, err := h.RepsitoryUpdatePromos(&body, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
		"message": "update promo success",
	})
}

func (h *HandlerPromos) DeletePromos(ctx *gin.Context) {
	id := ctx.Param("id")
	result, err := h.RepositoryDeletePromos(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		fmt.Println(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
		"message": "delete promo success",
	})
}