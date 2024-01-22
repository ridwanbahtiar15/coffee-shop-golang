package repositories

import (
	"coffee-shop-golang/internal/models"

	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	mock.Mock
}
// RepositoryGetAllProducts implements IProductsRepository.
func (prm *ProductRepositoryMock) RepositoryGetAllProducts(name string, category string, minrange string, maxrange string, page string, limit string, sort string) ([]models.ProductsResponseModel, error) {
	args := prm.Mock.Called(name, category, minrange, maxrange, page, limit, sort)
	return args.Get(0).([]models.ProductsResponseModel), args.Error(1)
}

// RepositoryCountProducts implements IProductsRepository.
func (prm *ProductRepositoryMock) RepositoryCountProducts(name string, category string, minrange string, maxrange string) ([]string, error) {
	args := prm.Mock.Called(name, category, minrange, maxrange)
	return args.Get(0).([]string), args.Error(1)
}

func (arm *ProductRepositoryMock) RepositoryCreateProducts(body *models.ProductsModel) ([]int, error) {
	args := arm.Mock.Called(body)
	return args.Get(0).([]int), args.Error(1)
}

// RepositoryDeleteProducts implements IProductsRepository.
func (prm *ProductRepositoryMock) RepositoryDeleteProducts(id string) (int64, error) {
	args := prm.Mock.Called(id)
	return args.Get(0).(int64), args.Error(1)
}

// RepositoryProductsById implements IProductsRepository.
func (prm *ProductRepositoryMock) RepositoryProductsById(id string) ([]models.ProductsResponseModel, error) {
	args := prm.Mock.Called(id)
	return args.Get(0).([]models.ProductsResponseModel), args.Error(1)
}

// RepositoryUpdateImgProducts implements IProductsRepository.
func (prm *ProductRepositoryMock) RepositoryUpdateImgProducts(productImg string, id string) error {
	panic("unimplemented")
}

// RepositoryUpdateProducts implements IProductsRepository.
func (prm *ProductRepositoryMock) RepositoryUpdateProducts(body *models.UpdateProductsModel, id string) error {
	panic("unimplemented")
}

// func (crm *ProductRepositoryMock) RepositoryGetAllProducts() ([]models.ProductsResponseModel, error) {
// 	// args := crm.Mock.Called(page, limit)
// 	// return args.Get(0).([]models.ProductsResponseModel), args.Error(1)
// }
