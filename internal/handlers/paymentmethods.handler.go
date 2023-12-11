package handlers

import (
	"coffee-shop-golang/internal/helpers"
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

	if len(result) == 0 {
		ctx.JSON(http.StatusNotFound, helpers.GetResponse("payment method not found", nil, nil))
		return
	}

	ctx.JSON(http.StatusOK, helpers.GetResponse("get all payment method success", nil, nil))
}