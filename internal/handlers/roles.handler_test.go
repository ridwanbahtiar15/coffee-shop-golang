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

var rrm = repositories.RoleRepositoryMock{}
var handlerRole = InitializeHandlerRoles(&rrm)

func TestGetAllRole(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	res := httptest.NewRecorder()
	r.GET("/roles", handlerRole.GetAllRoles)

	t.Run("Get Role", func(t *testing.T){
		ex := make([]models.RolesModel, 1)

		rrm.On("RepositoryGetAllRoles").Return(ex, nil)

		req := httptest.NewRequest("GET", "/roles", nil)
		r.ServeHTTP(res, req)

		exRes := helpers.GetResponse("get all role success", ex, nil)
		b, err := json.Marshal(exRes)
		if err != nil {
			t.Fatalf("Marshal Error: %e", err)
		}

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, string(b), res.Body.String())
	})
	
	t.Run("Role not found", func(t *testing.T){
		rrm = repositories.RoleRepositoryMock{}
		res := httptest.NewRecorder()
		ex := []models.RolesModel{}
		rrm.On("RepositoryGetAllRoles").Return(ex, nil)
 
		req := httptest.NewRequest("GET", "/roles", nil)
		r.ServeHTTP(res, req)

		exRes := helpers.GetResponse("role not found", nil, nil)
		b, err := json.Marshal(exRes)
		if err != nil {
			t.Fatalf("Marshal Error: %e", err)
		}

		assert.Equal(t, http.StatusNotFound, res.Code)
		assert.Equal(t, string(b), res.Body.String())
	})
}