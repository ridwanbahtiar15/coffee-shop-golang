package repositories

import (
	"coffee-shop-golang/internal/models"

	"github.com/stretchr/testify/mock"
)

type PromoRepositoryMock struct {
	mock.Mock
}

// RepositoryCountPromos implements IPromosRepository.
func (arm *PromoRepositoryMock) RepositoryCountPromos() ([]string, error) {
	args := arm.Mock.Called()
	return args.Get(0).([]string), args.Error(1)
}

// RepositoryDeletePromos implements IPromosRepository.
func (arm *PromoRepositoryMock) RepositoryDeletePromos(id string) (int64, error) {
	args := arm.Mock.Called(id)
	return args.Get(0).(int64), args.Error(1)
}

// RepositoryGetPromosById implements IPromosRepository.
func (arm *PromoRepositoryMock) RepositoryGetPromosById(id string) ([]models.PromosModel, error) {
	args := arm.Mock.Called(id)
	return args.Get(0).([]models.PromosModel), args.Error(1)
}

// RepsitoryCreatePromos implements IPromosRepository.
func (arm *PromoRepositoryMock) RepositoryCreatePromos(body *models.PromosModel) error {
	args := arm.Mock.Called(body)
	return args.Error(0)
}

// RepsitoryGetAllPromos implements IPromosRepository.
func (arm *PromoRepositoryMock) RepositoryGetAllPromos(page string, limit string) ([]models.PromosModel, error) {
	args := arm.Mock.Called(page, limit)
	return args.Get(0).([]models.PromosModel), args.Error(1)
}

// RepsitoryUpdatePromos implements IPromosRepository.
func (arm *PromoRepositoryMock) RepositoryUpdatePromos(body *models.UpdatePromosModel, id string) error {
	args := arm.Mock.Called(body, id)
	return args.Error(0)
}
