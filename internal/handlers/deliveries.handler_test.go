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

var drm = repositories.DeliveryRepositoryMock{}
var handlerDelivery = InitializeHandlerDeliveries(&drm)

func TestGetAllDelivery(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	res := httptest.NewRecorder()
	r.GET("/deliveries", handlerDelivery.GetAllDeliveries)

	t.Run("Get delivery", func(t *testing.T){
		ex := make([]models.DeliveriesModel, 1)

		drm.On("RepositoryGetAllDeliveries").Return(ex, nil)

		req := httptest.NewRequest("GET", "/deliveries", nil)
		r.ServeHTTP(res, req)

		exRes := helpers.GetResponse("get all delivery success", ex, nil)
		b, err := json.Marshal(exRes)
		if err != nil {
			t.Fatalf("Marshal Error: %e", err)
		}

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, string(b), res.Body.String())
	})
	
	t.Run("Delivery not found", func(t *testing.T){
		drm = repositories.DeliveryRepositoryMock{}
		res := httptest.NewRecorder()
		ex := []models.DeliveriesModel{}
		drm.On("RepositoryGetAllDeliveries").Return(ex, nil)
 
		req := httptest.NewRequest("GET", "/deliveries", nil)
		r.ServeHTTP(res, req)

		exRes := helpers.GetResponse("delivery not found", nil, nil)
		b, err := json.Marshal(exRes)
		if err != nil {
			t.Fatalf("Marshal Error: %e", err)
		}

		assert.Equal(t, http.StatusNotFound, res.Code)
		assert.Equal(t, string(b), res.Body.String())
	})
}