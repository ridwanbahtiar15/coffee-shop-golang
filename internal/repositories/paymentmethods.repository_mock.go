package repositories

import (
	"coffee-shop-golang/internal/models"

	"github.com/stretchr/testify/mock"
)

type PaymentMethodRepositoryMock struct {
	mock.Mock
}
func (crm *PaymentMethodRepositoryMock) RepositoryGetAllPaymentmethods() ([]models.PaymentmethodsModel, error) {
	args := crm.Mock.Called()
	return args.Get(0).([]models.PaymentmethodsModel), args.Error(1)
}
