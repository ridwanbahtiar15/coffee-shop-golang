package repositories

import (
	"coffee-shop-golang/internal/models"

	"github.com/jmoiron/sqlx"
)

type PaymentmethodsRepository struct {
	*sqlx.DB
}

func InitializeRepoPaymentmethods(db *sqlx.DB) *PaymentmethodsRepository {
	dr := PaymentmethodsRepository{db}
	return &dr
}

func (r *PaymentmethodsRepository) RepsitoryGetAllPaymentmethods() ([]models.PaymentmethodsModel, error) {
	result := []models.PaymentmethodsModel{}
	query := `SELECT * FROM payment_methods`
	err := r.Select(&result, query)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// func (r *DeliveriesRepository) RepsitoryCreateDeliveries(body *models.DeliveriesModel) (sql.Result, error) {
// 	query := `INSERT INTO deliveries (deliveries_name, deliveries_cost) VALUES (:deliveries_name, :deliveries_cost)`
// 	result, err := r.NamedExec(query, body)
// 	if err != nil {
// 		return result, err 
// 	}
// 	return result, nil
// }

// func (r *DeliveriesRepository) RepsitoryUpdateDeliveries(body *models.DeliveriesModel, id string) (sql.Result, error) {
// 	query := `UPDATE deliveries SET deliveries_name=:deliveries_name, deliveries_cost=:deliveries_cost WHERE deliveries_id =` + id
// 	result, err := r.NamedExec(query, body)
// 	if err != nil {
// 		return result, err 
// 	}
// 	return result, nil
// }

// func (r *DeliveriesRepository) RepositoryDeleteDeliveries(id string) (sql.Result, error) {
// 	query := `DELETE FROM deliveries WHERE deliveries_id = $1`
// 	result, err := r.Exec(query, id)
// 	if err != nil {
// 		return result, err 
// 	}
// 	return result, nil
// }