package handlers

import (
	"coffee-shop-golang/internal/helpers"
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

	if len(result) == 0 {
		ctx.JSON(http.StatusNotFound, helpers.GetResponse("delivery not found", nil))
		return
	}

	ctx.JSON(http.StatusOK, helpers.GetResponse("get all delivery success", result))
}