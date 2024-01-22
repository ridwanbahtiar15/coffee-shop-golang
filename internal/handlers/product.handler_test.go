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

var prm = repositories.ProductRepositoryMock{}
var handler = InitializeHandlerProducts(&prm)

func TestGetAllProduct(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	res := httptest.NewRecorder()
	r.GET("/products", handler.GetAllProducts)

	t.Run("Get product", func(t *testing.T) {
		count := make([]string, 2)
		ex := make([]models.ProductsResponseModel, 2)
		meta := &helpers.Meta{
			Page:     1,
			TotalData: 2,
			NextPage: "null",
			PrevPage: "null",
		}
		
		prm.On("RepositoryCountProducts", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(count, nil)

		prm.On("RepositoryGetAllProducts", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(ex, nil)

		req := httptest.NewRequest("GET", "/products", nil)
		r.ServeHTTP(res, req)

		exRes := helpers.GetResponse("get product success", ex, meta)
		b, err := json.Marshal(exRes)
		if err != nil {
			t.Fatalf("marshal error : %e", err)
		}
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, string(b), res.Body.String())
	})

	t.Run("Invalid minrange and maxrange", func(t *testing.T) {
		res = httptest.NewRecorder()
		count := make([]string, 2)
		ex := make([]models.ProductsResponseModel, 2)
		
		prm.On("RepositoryCountProducts", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(count, nil)

		prm.On("RepositoryGetAllProducts", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(ex, nil)

		req := httptest.NewRequest("GET", "/products?minrange=50000&maxrange=30000", nil)
		r.ServeHTTP(res, req)

		exRes := helpers.GetResponse("The range your input is not correct", nil, nil)
		b, err := json.Marshal(exRes)
		if err != nil {
			t.Fatalf("marshal error : %e", err)
		}
		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, string(b), res.Body.String())
	})

	t.Run("Product not found", func(t *testing.T) {
		prm = repositories.ProductRepositoryMock{}
		res = httptest.NewRecorder()
		ex := []models.ProductsResponseModel{}
		prm.On("RepositoryGetAllProducts", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(ex, nil)

		req := httptest.NewRequest("GET", "/products", nil)
		r.ServeHTTP(res, req)

		exRes := helpers.GetResponse("product not found", nil, nil)
		b, err := json.Marshal(exRes)
		if err != nil {
			t.Fatalf("marshal error : %e", err)
		}
		assert.Equal(t, http.StatusNotFound, res.Code)
		assert.Equal(t, string(b), res.Body.String())
	})
}

func TestGetProductById(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	res := httptest.NewRecorder()
	r.GET("/products/39", handler.GetProductsById)

	t.Run("Get product by id success", func(t *testing.T) {
		ex := make([]models.ProductsResponseModel, 1)

		prm.On("RepositoryProductsById", mock.AnythingOfType("string")).Return(ex, nil)

		req := httptest.NewRequest("GET", "/products/39", nil)
		r.ServeHTTP(res, req)

		exRes := helpers.GetResponse("get product by id success", ex, nil)
		b, err := json.Marshal(exRes)
		if err != nil {
			t.Fatalf("marshal error : %e", err)
		}
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, string(b), res.Body.String())
	})

	t.Run("Get product by id not found", func(t *testing.T) {
		prm = repositories.ProductRepositoryMock{}
		res = httptest.NewRecorder()
		ex := make([]models.ProductsResponseModel, 0)

		prm.On("RepositoryProductsById", mock.AnythingOfType("string")).Return(ex, nil)

		req := httptest.NewRequest("GET", "/products/39", nil)
		r.ServeHTTP(res, req)

		exRes := helpers.GetResponse("product not found", nil, nil)
		b, err := json.Marshal(exRes)
		if err != nil {
			t.Fatalf("marshal error : %e", err)
		}
		assert.Equal(t, http.StatusNotFound, res.Code)
		assert.Equal(t, string(b), res.Body.String())
	})
}

func TestCreatePoduct(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	res := httptest.NewRecorder()

	r.POST("/products", handler.CreateProducts)
		prm.On("RepositoryCreateProducts", mock.Anything).Return([]int{}, nil)
	
		body := &models.ProductsModel{
			Products_name:    "Kopi ABC",
			Products_price:   "32000",
			Products_desc:    "ini kopi abc",
			Products_stock: 	"100",
			Categories_id:    "1",
			Products_image: 	"kopi.jpg",
		}
		b, err := json.Marshal(body)
		if err != nil {
			t.Fatalf("Marshal Error: %e", err)
		}

	req := httptest.NewRequest("POST", "/products", strings.NewReader(string(b)))
	req.Header.Set("Content-Type", "multipart/form-data")
	r.ServeHTTP(res, req)

	exRes := helpers.GetResponse("create product success", nil, nil)
	bres, err := json.Marshal(exRes)
	if err != nil {
		t.Fatalf("Marshal Error: %e", err)
	}

	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, string(bres), res.Body.String())
}

func TestDeleteProduct(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	res := httptest.NewRecorder()
	r.DELETE("/products/2", handler.DeleteProducts)

	t.Run("Delete product success", func(t *testing.T) {
		var data int64 = 1
		prm.On("RepositoryDeleteProducts", mock.AnythingOfType("string")).Return(data, nil)

		req := httptest.NewRequest("DELETE", "/products/2", nil)
		r.ServeHTTP(res, req)

		exRes := helpers.GetResponse("delete product success", nil, nil)
		b, err := json.Marshal(exRes)
		if err != nil {
			t.Fatalf("marshal error : %e", err)
		}
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, string(b), res.Body.String())
	})

	t.Run("Delete product id not found", func(t *testing.T) {
		prm = repositories.ProductRepositoryMock{}
		res = httptest.NewRecorder()
		var data int64 = 0
		prm.On("RepositoryDeleteProducts", mock.AnythingOfType("string")).Return(data, nil)

		req := httptest.NewRequest("DELETE", "/products/2", nil)
		r.ServeHTTP(res, req)

		exRes := helpers.GetResponse("id product not found", nil, nil)
		b, err := json.Marshal(exRes)
		if err != nil {
			t.Fatalf("marshal error : %e", err)
		}
		assert.Equal(t, http.StatusNotFound, res.Code)
		assert.Equal(t, string(b), res.Body.String())
	})
}