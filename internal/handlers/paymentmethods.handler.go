package handlers

import (
	"coffee-shop-golang/internal/helpers"
	"coffee-shop-golang/internal/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerPaymentmethods struct {
	repositories.IPaymentMethodsRepository
}

func InitializeHandlerPaymentmethods(r repositories.IPaymentMethodsRepository) *HandlerPaymentmethods {
	return &HandlerPaymentmethods{r}
}

func (h *HandlerPaymentmethods) GetAllPaymentmethods(ctx *gin.Context) {
	result, err := h.RepositoryGetAllPaymentmethods()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if len(result) == 0 {
		ctx.JSON(http.StatusNotFound, helpers.GetResponse("payment method not found", nil, nil))
		return
	}

	ctx.JSON(http.StatusOK, helpers.GetResponse("get all payment method success", result, nil))
}