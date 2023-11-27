package handlers

import (
	"coffee-shop-golang/internal/models"
	"coffee-shop-golang/internal/repositories"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerOrders struct {
	*repositories.OrdersRepository
}

func InitializeHandlerOrders(r *repositories.OrdersRepository) *HandlerOrders {
	return &HandlerOrders{r}
}

func (h *HandlerOrders) GetAllOrders(ctx *gin.Context) {
	result, err := h.RepsitoryGetAllOrders()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if len(result) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "order not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
		"message": "get all order success",
	})
}

func (h *HandlerOrders) CreateOrders(ctx *gin.Context) {
	var body models.OrdersModel
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		fmt.Println(err)
	}

	err := h.RepsitoryCreateOrders(&body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "create order success",
	})
}