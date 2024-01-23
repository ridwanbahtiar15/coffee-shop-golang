package handlers

import (
	"coffee-shop-golang/internal/helpers"
	"coffee-shop-golang/internal/models"
	"coffee-shop-golang/internal/repositories"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var orm = repositories.OrderRepositoryMock{}
var handlerOrder = InitializeHandlerOrders(&orm)

func TestGetAllOrders(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	res := httptest.NewRecorder()
	r.GET("/orders", handlerOrder.GetAllOrders)

	t.Run("Get order", func(t *testing.T){
		ex := make([]models.OrdersResponseModel, 1)
		meta := &helpers.Meta{
			Page: 1,
			TotalData: 1,
			NextPage: "null",
			PrevPage: "null",
		}

		var count = []string{"1"}
		orm.On("RepositoryCountOrders", mock.AnythingOfType("string")).Return(count, nil)
		orm.On("RepositoryGetAllOrders", mock.AnythingOfType("string"), mock.AnythingOfType("string"),  mock.AnythingOfType("string"),  mock.AnythingOfType("string")).Return(ex, nil)

		req := httptest.NewRequest("GET", "/orders", nil)
		r.ServeHTTP(res, req)

		exRes := helpers.GetResponse("get order success", ex, meta)
		b, err := json.Marshal(exRes)
		if err != nil {
			t.Fatalf("Marshal Error: %e", err)
		}

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, string(b), res.Body.String())
	})
	
	t.Run("Order not found", func(t *testing.T){
		orm = repositories.OrderRepositoryMock{}
		res := httptest.NewRecorder()
		ex := []models.OrdersResponseModel{}
		orm.On("RepositoryGetAllOrders", mock.AnythingOfType("string"), mock.AnythingOfType("string"),  mock.AnythingOfType("string"),  mock.AnythingOfType("string")).Return(ex, nil)

		req := httptest.NewRequest("GET", "/orders", nil)
		r.ServeHTTP(res, req)

		exRes := helpers.GetResponse("order not found", nil, nil)
		b, err := json.Marshal(exRes)
		if err != nil {
			t.Fatalf("Marshal Error: %e", err)
		}

		assert.Equal(t, http.StatusNotFound, res.Code)
		assert.Equal(t, string(b), res.Body.String())
	})
}

func TestOrdersById(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	res := httptest.NewRecorder()
	r.GET("/orders/1", handlerOrder.GetOrdersById)

	t.Run("Get order by id", func(t *testing.T){
		ex := make([]models.OrdersResponseModel, 1)
		orm.On("RepositoryGetOrdersById", mock.AnythingOfType("string")).Return(ex, nil)

		req := httptest.NewRequest("GET", "/orders/1", nil)
		r.ServeHTTP(res, req)

		exRes := helpers.GetResponse("get order by id success", ex, nil)
		b, err := json.Marshal(exRes)
		if err != nil {
			t.Fatalf("Marshal Error: %e", err)
		}

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, string(b), res.Body.String())
	})

	t.Run("Order by id not found", func(t *testing.T){
		orm = repositories.OrderRepositoryMock{}
		res := httptest.NewRecorder()
		ex := []models.OrdersResponseModel{}
		orm.On("RepositoryGetOrdersById", mock.AnythingOfType("string")).Return(ex, nil)

		req := httptest.NewRequest("GET", "/orders/1", nil)
		r.ServeHTTP(res, req)

		exRes := helpers.GetResponse("order not found", nil, nil)
		b, err := json.Marshal(exRes)
		if err != nil {
			t.Fatalf("Marshal Error: %e", err)
		}

		assert.Equal(t, http.StatusNotFound, res.Code)
		assert.Equal(t, string(b), res.Body.String())
	})
}

func TestCreateOrder(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	res := httptest.NewRecorder()
	r.POST("/orders", handlerOrder.CreateOrders)

	t.Run("create order success", func(t *testing.T) {
		orm.On("RepositoryCreateOrders", mock.Anything).Return(nil)

		body := &models.OrdersModel{
			Users_id :    "1",
			Deliveries_id:   "1",
			Promos_id:    "1",
			Payment_methods_id: "2",
			Orders_status: "Pending",
			Orders_total: "10000",
			Orders_products_id: "1",
			Products_id: "1",
			Sizes_id: "2",
			Orders_products_qty: "2",
			Orders_products_subtotal: "10000",
			Hot_or_ice: "Hot",
		}
		b, err := json.Marshal(body)
		if err != nil {
			t.Fatalf("Marshal Error: %e", err)
		}

		req := httptest.NewRequest("POST", "/orders", strings.NewReader(string(b)))
		req.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(res, req)

		exRes := helpers.GetResponse("create order success", nil, nil)
		bres, err := json.Marshal(exRes)
		if err != nil {
			t.Fatalf("Marshal Error: %e", err)
		}

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, string(bres), res.Body.String())
	})
}

func TestUpdateOrders(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	res := httptest.NewRecorder()

	r.PATCH("/orders/1", handlerOrder.UpdateOrders)
	t.Run("Update order success", func(t *testing.T) {
		ex := make([]models.OrdersModel, 1)
		orm.On("RepositoryGetOrdersById", mock.AnythingOfType("string")).Return(ex, nil)
		orm.On("RepositoryUpdateOrders", mock.Anything, mock.AnythingOfType("string")).Return(nil)
		
		body := &models.OrdersModel{
			Orders_status :    "done",
		}
		b, err := json.Marshal(body)
		if err != nil {
			t.Fatalf("Marshal Error: %e", err)
		}

		req := httptest.NewRequest("PATCH", "/orders/1", strings.NewReader(string(b)))
		req.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(res, req)

		exRes := helpers.GetResponse("update order success", nil, nil)
		bres, err := json.Marshal(exRes)
		if err != nil {
			t.Fatalf("Marshal Error: %e", err)
		}

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, string(bres), res.Body.String())
		})
}
