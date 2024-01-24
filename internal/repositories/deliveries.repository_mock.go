package repositories

import (
	"coffee-shop-golang/internal/models"

	"github.com/stretchr/testify/mock"
)

type DeliveryRepositoryMock struct {
	mock.Mock
}
func (crm *DeliveryRepositoryMock) RepositoryGetAllDeliveries() ([]models.DeliveriesModel, error) {
	args := crm.Mock.Called()
	return args.Get(0).([]models.DeliveriesModel), args.Error(1)
}
