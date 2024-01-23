package repositories

import (
	"coffee-shop-golang/internal/models"

	"github.com/stretchr/testify/mock"
)

type OrderRepositoryMock struct {
	mock.Mock
}

// RepsitoryGetAllPromos implements IPromosRepository.
func (orm *OrderRepositoryMock) RepositoryGetAllOrders(orderNumber string, page string, limit string, sort string) ([]models.OrdersResponseModel, error) {
	args := orm.Mock.Called(orderNumber, page, limit, sort)
	return args.Get(0).([]models.OrdersResponseModel), args.Error(1)
}

func (orm *OrderRepositoryMock) RepositoryGetOrdersById(id string) ([]models.OrdersResponseModel, error) {
	args := orm.Mock.Called(id)
	return args.Get(0).([]models.OrdersResponseModel), args.Error(1)
}

// RepositoryCountPromos implements IPromosRepository.
func (orm *OrderRepositoryMock) RepositoryCountOrders(orderNumber string) ([]string, error) {
	args := orm.Mock.Called(orderNumber)
	return args.Get(0).([]string), args.Error(1)
}

// RepositoryGetPromosById implements IPromosRepository.
// func (orm *OrderRepositoryMock) RepositoryGetPromosById(id string) ([]models.PromosModel, error) {
// 	args := orm.Mock.Called(id)
// 	return args.Get(0).([]models.PromosModel), args.Error(1)
// }

// RepsitoryCreatePromos implements IPromosRepository.
func (orm *OrderRepositoryMock) RepositoryCreateOrders(body *models.OrdersModel) (error) {
	args := orm.Mock.Called(body)
	return args.Error(0)
}


// RepsitoryUpdatePromos implements IPromosRepository.
func (orm *OrderRepositoryMock) RepositoryUpdateOrders(body *models.OrderUpdateModel, id string) (error) {
	args := orm.Mock.Called(body, id)
	return args.Error(0)
}
