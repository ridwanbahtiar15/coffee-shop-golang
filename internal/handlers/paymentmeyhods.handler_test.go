package handlers

import (
	"coffee-shop-golang/internal/helpers"
	"coffee-shop-golang/internal/models"
	"coffee-shop-golang/internal/repositories"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var pmm = repositories.PaymentMethodRepositoryMock{}
var handlerPaymentMethod = InitializeHandlerPaymentmethods(&pmm)

func TestGetAllPaymentMethod(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	res := httptest.NewRecorder()
	r.GET("/paymentmethods", handlerPaymentMethod.GetAllPaymentmethods)

	t.Run("Get payment method", func(t *testing.T){
		ex := make([]models.PaymentmethodsModel, 1)

		pmm.On("RepositoryGetAllPaymentmethods").Return(ex, nil)

		req := httptest.NewRequest("GET", "/paymentmethods", nil)
		r.ServeHTTP(res, req)

		exRes := helpers.GetResponse("get all payment method success", ex, nil)
		b, err := json.Marshal(exRes)
		if err != nil {
			t.Fatalf("Marshal Error: %e", err)
		}

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, string(b), res.Body.String())
	})
	
	t.Run("Payment method not found", func(t *testing.T){
		pmm = repositories.PaymentMethodRepositoryMock{}
		res := httptest.NewRecorder()
		ex := []models.PaymentmethodsModel{}
		pmm.On("RepositoryGetAllPaymentmethods").Return(ex, nil)
 
		req := httptest.NewRequest("GET", "/paymentmethods", nil)
		r.ServeHTTP(res, req)

		exRes := helpers.GetResponse("payment method not found", nil, nil)
		b, err := json.Marshal(exRes)
		if err != nil {
			t.Fatalf("Marshal Error: %e", err)
		}

		assert.Equal(t, http.StatusNotFound, res.Code)
		assert.Equal(t, string(b), res.Body.String())
	})
}