package repositories

import (
	"coffee-shop-golang/internal/models"

	"github.com/jmoiron/sqlx"
)

type IPaymentMethodsRepository interface {
	RepositoryGetAllPaymentmethods() ([]models.PaymentmethodsModel, error)
}

type PaymentmethodsRepository struct {
	*sqlx.DB
}

func InitializeRepoPaymentmethods(db *sqlx.DB) *PaymentmethodsRepository {
	dr := PaymentmethodsRepository{db}
	return &dr
}

func (r *PaymentmethodsRepository) RepositoryGetAllPaymentmethods() ([]models.PaymentmethodsModel, error) {
	result := []models.PaymentmethodsModel{}
	query := `SELECT payment_methods_id, payment_methods_name FROM payment_methods`
	err := r.Select(&result, query)
	if err != nil {
		return nil, err
	}
	return result, nil
}