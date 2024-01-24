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

var crm = repositories.CategoryRepositoryMock{}
var handlerCategory = InitializeHandlerCategories(&crm)

func TestGetAllCategory(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	res := httptest.NewRecorder()
	r.GET("/categories", handlerCategory.GetAllCategories)

	t.Run("Get category", func(t *testing.T){
		ex := make([]models.CategoriesModel, 1)

		crm.On("RepositoryGetAllCategories").Return(ex, nil)

		req := httptest.NewRequest("GET", "/categories", nil)
		r.ServeHTTP(res, req)

		exRes := helpers.GetResponse("get all category success", ex, nil)
		b, err := json.Marshal(exRes)
		if err != nil {
			t.Fatalf("Marshal Error: %e", err)
		}

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, string(b), res.Body.String())
	})
	
	t.Run("Category not found", func(t *testing.T){
		crm = repositories.CategoryRepositoryMock{}
		res := httptest.NewRecorder()
		ex := []models.CategoriesModel{}
		crm.On("RepositoryGetAllCategories").Return(ex, nil)
 
		req := httptest.NewRequest("GET", "/categories", nil)
		r.ServeHTTP(res, req)

		exRes := helpers.GetResponse("category not found", nil, nil)
		b, err := json.Marshal(exRes)
		if err != nil {
			t.Fatalf("Marshal Error: %e", err)
		}

		assert.Equal(t, http.StatusNotFound, res.Code)
		assert.Equal(t, string(b), res.Body.String())
	})
}