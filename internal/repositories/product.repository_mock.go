package repositories

import (
	"coffee-shop-golang/internal/models"

	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	mock.Mock
}

func (crm *ProductRepositoryMock) GetAllProduct(page string, limit string) ([]models.ProductsResponseModel, error) {
	args := crm.Mock.Called(page, limit)
	return args.Get(0).([]models.ProductsResponseModel), args.Error(1)
}