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

var arm = repositories.PromoRepositoryMock{}
var handlerPromo = InitializeHandlerPromos(&arm)

func TestGetAllPromos(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	res := httptest.NewRecorder()
	r.GET("/promos", handlerPromo.GetAllPromos)

	t.Run("Get promo", func(t *testing.T){
		ex := make([]models.PromosModel, 1)
		meta := &helpers.Meta{
			TotalData: 1,
			NextPage: "null",
			PrevPage: "null",
		}

		var count = []string{"1"}
		arm.On("RepositoryCountPromos").Return(count, nil)
		arm.On("RepositoryGetAllPromos", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(ex, nil)

		req := httptest.NewRequest("GET", "/promos", nil)
		r.ServeHTTP(res, req)

		exRes := helpers.GetResponse("get promo success", ex, meta)
		b, err := json.Marshal(exRes)
		if err != nil {
			t.Fatalf("Marshal Error: %e", err)
		}

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, string(b), res.Body.String())
	})
	
	t.Run("Promo not found", func(t *testing.T){
		arm = repositories.PromoRepositoryMock{}
		res := httptest.NewRecorder()
		ex := []models.PromosModel{}
		arm.On("RepositoryGetAllPromos", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(ex, nil)

		req := httptest.NewRequest("GET", "/promos", nil)
		r.ServeHTTP(res, req)

		exRes := helpers.GetResponse("promo not found", nil, nil)
		b, err := json.Marshal(exRes)
		if err != nil {
			t.Fatalf("Marshal Error: %e", err)
		}

		assert.Equal(t, http.StatusNotFound, res.Code)
		assert.Equal(t, string(b), res.Body.String())
	})

	// t.Run("Internal server error", func(t *testing.T){
	// 	ex := []models.PromosModel{}
	// 	arm.On("RepositoryGetAllPromos", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(ex, errors.New("Internal Server Error"))

	// 	req := httptest.NewRequest("GET", "/promos", nil)
	// 	r.ServeHTTP(res, req)

	// 	exRes := helpers.GetResponse("Internal Server Error", ex, nil)
	// 	b, err := json.Marshal(exRes)
	// 	if err != nil {
	// 		t.Fatalf("Marshal Error: %e", err)
	// 	}

	// 	assert.Equal(t, http.StatusInternalServerError, res.Code)
	// 	assert.Equal(t, string(b), res.Body.String())
	// })
}

func TestPromosById(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	res := httptest.NewRecorder()
	r.GET("/promos/4", handlerPromo.GetPromosById)

	t.Run("Get promo by id", func(t *testing.T){
		ex := make([]models.PromosModel, 1)
		arm.On("RepositoryGetPromosById", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(ex, nil)

		req := httptest.NewRequest("GET", "/promos/4", nil)
		r.ServeHTTP(res, req)

		exRes := helpers.GetResponse("get promo by id success", ex, nil)
		b, err := json.Marshal(exRes)
		if err != nil {
			t.Fatalf("Marshal Error: %e", err)
		}

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, string(b), res.Body.String())
	})

	t.Run("Promo by id not found", func(t *testing.T){
		arm = repositories.PromoRepositoryMock{}
		res := httptest.NewRecorder()
		ex := []models.PromosModel{}
		arm.On("RepositoryGetPromosById", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(ex, nil)

		req := httptest.NewRequest("GET", "/promos/4", nil)
		r.ServeHTTP(res, req)

		exRes := helpers.GetResponse("promo not found", nil, nil)
		b, err := json.Marshal(exRes)
		if err != nil {
			t.Fatalf("Marshal Error: %e", err)
		}

		assert.Equal(t, http.StatusNotFound, res.Code)
		assert.Equal(t, string(b), res.Body.String())
	})
}

func TestCreatePromos(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	res := httptest.NewRecorder()
	r.POST("/promos", handlerPromo.CreateProomos)

	t.Run("create promo success", func(t *testing.T) {
		arm.On("RepositoryCreatePromos", mock.Anything).Return(nil)

		body := &models.PromosModel{
			Promos_name:    "PROMO1212",
			Promos_start:   "12-12-2023",
			Promos_end:    "13-12-2023",
		}
		b, err := json.Marshal(body)
		if err != nil {
			t.Fatalf("Marshal Error: %e", err)
		}

		req := httptest.NewRequest("POST", "/promos", strings.NewReader(string(b)))
		req.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(res, req)

		exRes := helpers.GetResponse("create promo success", nil, nil)
		bres, err := json.Marshal(exRes)
		if err != nil {
			t.Fatalf("Marshal Error: %e", err)
		}

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, string(bres), res.Body.String())
	})
}

func TestUpdatePromos(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	res := httptest.NewRecorder()

	r.PATCH("/promos/1", handlerPromo.UpdatePromos)
	t.Run("Update promo success", func(t *testing.T) {
		ex := make([]models.PromosModel, 1)
		arm.On("RepositoryGetPromosById", mock.AnythingOfType("string")).Return(ex, nil)
		arm.On("RepositoryUpdatePromos", mock.Anything, mock.AnythingOfType("string")).Return(nil)
		
		body := &models.PromosModel{
			Promos_name:    "PROMO1212",
			Promos_start:   "12-12-2023",
			Promos_end:    "13-12-2023",
		}
		b, err := json.Marshal(body)
		if err != nil {
			t.Fatalf("Marshal Error: %e", err)
		}

		req := httptest.NewRequest("PATCH", "/promos/1", strings.NewReader(string(b)))
		req.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(res, req)

		exRes := helpers.GetResponse("update promo success", nil, nil)
		bres, err := json.Marshal(exRes)
		if err != nil {
			t.Fatalf("Marshal Error: %e", err)
		}

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, string(bres), res.Body.String())
		})
}

func TestDeletePromo(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	res := httptest.NewRecorder()
	r.DELETE("/promos/17", handlerPromo.DeletePromos)

	t.Run("Delete promo success", func(t *testing.T) {
		var data int64 = 1
		arm.On("RepositoryDeletePromos", mock.AnythingOfType("string")).Return(data, nil)

		req := httptest.NewRequest("DELETE", "/promos/17", nil)
		r.ServeHTTP(res, req)

		exRes := helpers.GetResponse("delete promo success", nil, nil)
		b, err := json.Marshal(exRes)
		if err != nil {
			t.Fatalf("marshal error : %e", err)
		}
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, string(b), res.Body.String())
	})

	t.Run("Delete promo id not found", func(t *testing.T) {
		arm = repositories.PromoRepositoryMock{}
		res = httptest.NewRecorder()
		var data int64 = 0
		arm.On("RepositoryDeletePromos", mock.AnythingOfType("string")).Return(data, nil)

		req := httptest.NewRequest("DELETE", "/promos/17", nil)
		r.ServeHTTP(res, req)

		exRes := helpers.GetResponse("id promo not found", nil, nil)
		b, err := json.Marshal(exRes)
		if err != nil {
			t.Fatalf("marshal error : %e", err)
		}
		assert.Equal(t, http.StatusNotFound, res.Code)
		assert.Equal(t, string(b), res.Body.String())
	})
}