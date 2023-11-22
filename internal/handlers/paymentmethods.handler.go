package handlers

import (
	"coffee-shop-golang/internal/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerPaymentmethods struct {
	*repositories.PaymentmethodsRepository
}

func InitializeHandlerPaymentmethods(r *repositories.PaymentmethodsRepository) *HandlerPaymentmethods {
	return &HandlerPaymentmethods{r}
}

func (h *HandlerPaymentmethods) GetAllPaymentmethods(ctx *gin.Context) {
	result, err := h.RepsitoryGetAllPaymentmethods()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
		"message": "get all payment method success",
	})
}