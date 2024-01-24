package repositories

import (
	"coffee-shop-golang/internal/models"

	"github.com/stretchr/testify/mock"
)

type CategoryRepositoryMock struct {
	mock.Mock
}

// RepsitoryGetAllCategories implements ICategoriesRepository.
func (crm *CategoryRepositoryMock) RepositoryGetAllCategories() ([]models.CategoriesModel, error) {
	args := crm.Mock.Called()
	return args.Get(0).([]models.CategoriesModel), args.Error(1)
}
