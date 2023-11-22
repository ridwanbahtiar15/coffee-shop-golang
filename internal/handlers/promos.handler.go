package handlers

import (
	"coffee-shop-golang/internal/repositories"
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