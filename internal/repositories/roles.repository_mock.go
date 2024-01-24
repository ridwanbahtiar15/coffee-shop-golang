package repositories

import (
	"coffee-shop-golang/internal/models"

	"github.com/stretchr/testify/mock"
)

type RoleRepositoryMock struct {
	mock.Mock
}
func (rrm *RoleRepositoryMock) RepositoryGetAllRoles() ([]models.RolesModel, error) {
	args := rrm.Mock.Called()
	return args.Get(0).([]models.RolesModel), args.Error(1)
}
