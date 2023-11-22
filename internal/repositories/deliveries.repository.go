package repositories

import (
	"coffee-shop-golang/internal/models"

	"github.com/jmoiron/sqlx"
)

type DeliveriesRepository struct {
	*sqlx.DB
}

func InitializeRepoDeliveries(db *sqlx.DB) *DeliveriesRepository {
	dr := DeliveriesRepository{db}
	return &dr
}

func (r *DeliveriesRepository) RepsitoryGetAllDeliveries() ([]models.DeliveriesModel, error) {
	result := []models.DeliveriesModel{}
	query := `SELECT * FROM deliveries`
	err := r.Select(&result, query)
	if err != nil {
		return nil, err
	}
	return result, nil
}