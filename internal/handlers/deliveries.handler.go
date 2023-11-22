package handlers

import (
	"coffee-shop-golang/internal/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerDeliveries struct {
	*repositories.DeliveriesRepository
}

func InitializeHandlerDeliveries(r *repositories.DeliveriesRepository) *HandlerDeliveries {
	return &HandlerDeliveries{r}
}

func (h *HandlerDeliveries) GetAllDeliveries(ctx *gin.Context) {
	result, err := h.RepsitoryGetAllDeliveries()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
		"message": "get all delivery success",
	})
}