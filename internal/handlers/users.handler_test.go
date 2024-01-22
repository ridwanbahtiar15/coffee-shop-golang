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

var urm = repositories.UsersRepositoryMock{}
var handlerUsers = InitializeHandlerUsers(&urm)

func TestGetAllUsers(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	res := httptest.NewRecorder()
	r.GET("/users", handlerUsers.GetAllUsers)

	t.Run("Get user", func(t *testing.T){
		ex := make([]models.UsersResponseModel, 2)
		meta := &helpers.Meta{
			Page:     1,
			NextPage: "null",
			PrevPage: "null",
		}

		urm.On("RepositoryCountUsers", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(make([]string, 2), nil)
		urm.On("RepositoryGetAllUsers", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(ex, nil)

		req := httptest.NewRequest("GET", "/users", nil)
		r.ServeHTTP(res, req)

		exRes := helpers.GetResponse("get user success", ex, meta)
		b, err := json.Marshal(exRes)
		if err != nil {
			t.Fatalf("Marshal Error: %e", err)
		}

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, string(b), res.Body.String())
	})
	
	t.Run("User not found", func(t *testing.T){
		urm = repositories.UsersRepositoryMock{}
		res := httptest.NewRecorder()
		ex := []models.UsersResponseModel{}
		urm.On("RepositoryGetAllUsers", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(ex, nil)

		req := httptest.NewRequest("GET", "/users", nil)
		r.ServeHTTP(res, req)

		exRes := helpers.GetResponse("user not found", nil, nil)
		b, err := json.Marshal(exRes)
		if err != nil {
			t.Fatalf("Marshal Error: %e", err)
		}

		assert.Equal(t, http.StatusNotFound, res.Code)
		assert.Equal(t, string(b), res.Body.String())
	})
}

func TestUserById(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	res := httptest.NewRecorder()
	r.GET("/users/4", handlerUsers.GetUsersById)

	t.Run("Get user by id", func(t *testing.T){
		ex := make([]models.UsersResponseModel, 1)
		urm.On("RepositoryUsersById", mock.AnythingOfType("string")).Return(ex, nil)

		req := httptest.NewRequest("GET", "/users/4", nil)
		r.ServeHTTP(res, req)

		exRes := helpers.GetResponse("get user by id success", ex, nil)
		b, err := json.Marshal(exRes)
		if err != nil {
			t.Fatalf("Marshal Error: %e", err)
		}

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, string(b), res.Body.String())
	})

	t.Run("Get not found user by id", func(t *testing.T){
		var urm = repositories.UsersRepositoryMock{}
		res := httptest.NewRecorder()
		ex := []models.UsersResponseModel{}
		urm.On("RepositoryUsersById", mock.AnythingOfType("string")).Return(ex, nil)

		req := httptest.NewRequest("GET", "/users/4", nil)
		r.ServeHTTP(res, req)

		exRes := helpers.GetResponse("user not found", nil, nil)
		b, err := json.Marshal(exRes)
		if err != nil {
			t.Fatalf("Marshal Error: %e", err)
		}

		assert.Equal(t, http.StatusNotFound, res.Code)
		assert.Equal(t, string(b), res.Body.String())
	})
}

func TestCreateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	res := httptest.NewRecorder()

	r.POST("/users", handlerUsers.CreateUsers)
	t.Run("create user success", func(t *testing.T) {
		var id int = 1

		urm.On("RepositoryCreateUsers", mock.Anything).Return(id, nil)
		urm.On("RepositoryUpdateImgUsers", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil)

		body := &models.UsersModel{
			Users_fullname:    "ridwan",
			Users_email:   "ridwan@mail.com",
			Users_password:    "12345",
			Users_phone: "08123232332",
			Users_address: "bekasi",
			Roles_id: "1",
		}
		b, err := json.Marshal(body)
		if err != nil {
			t.Fatalf("Marshal Error: %e", err)
		}

		req := httptest.NewRequest("POST", "/users", strings.NewReader(string(b)))
		req.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(res, req)

		exRes := helpers.GetResponse("create user success", nil, nil)
		bres, err := json.Marshal(exRes)
		if err != nil {
			t.Fatalf("Marshal Error: %e", err)
		}

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, string(bres), res.Body.String())
	})
}

func TestUpdateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	res := httptest.NewRecorder()

	r.PATCH("/users/1", handlerUsers.UpdateUsers)
	t.Run("update user success", func(t *testing.T) {
		var id int = 1
		ex := make([]models.UsersResponseModel, 1)
		urm.On("RepositoryUsersById", mock.AnythingOfType("string")).Return(ex, nil)
		urm.On("RepositoryUpdateUsers", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(id, nil)
		urm.On("RepositoryUpdateImgUsers", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil)

		body := &models.UsersModel{
			Users_fullname:    "ridwan",
			Users_email:   "ridwan@mail.com",
			Users_password:    "12345",
			Users_phone: "08123232332",
			Users_address: "bekasi",
			Roles_id: "1",
		}
		b, err := json.Marshal(body)
		if err != nil {
			t.Fatalf("Marshal Error: %e", err)
		}

		req := httptest.NewRequest("PATCH", "/users/1", strings.NewReader(string(b)))
		req.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(res, req)

		exRes := helpers.GetResponse("update user success", nil, nil)
		bres, err := json.Marshal(exRes)
		if err != nil {
			t.Fatalf("Marshal Error: %e", err)
		}

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, string(bres), res.Body.String())
	})
}

func TestDeleteUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	res := httptest.NewRecorder()
	r.DELETE("/users/1", handlerUsers.DeleteUsers)

	t.Run("Delete user success", func(t *testing.T) {
		var numb int = 1
		urm.On("RepositoryDeleteUsers", mock.AnythingOfType("string")).Return(numb, nil)

		req := httptest.NewRequest("DELETE", "/users/1", nil)
		r.ServeHTTP(res, req)

		exRes := helpers.GetResponse("delete user success", nil, nil)
		b, err := json.Marshal(exRes)
		if err != nil {
			t.Fatalf("marshal error : %e", err)
		}
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, string(b), res.Body.String())
	})

		t.Run("Delete user id not found", func(t *testing.T) {
			urm = repositories.UsersRepositoryMock{}
			res = httptest.NewRecorder()
			var numb int = 0
			urm.On("RepositoryDeleteUsers", mock.AnythingOfType("string")).Return(numb, nil)
	
			req := httptest.NewRequest("DELETE", "/users/1", nil)
			r.ServeHTTP(res, req)
	
			exRes := helpers.GetResponse("id user not found", nil, nil)
			b, err := json.Marshal(exRes)
			if err != nil {
				t.Fatalf("marshal error : %e", err)
			}
			assert.Equal(t, http.StatusNotFound, res.Code)
			assert.Equal(t, string(b), res.Body.String())
	})
}