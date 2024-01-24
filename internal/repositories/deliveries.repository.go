package repositories

import (
	"coffee-shop-golang/internal/models"

	"github.com/jmoiron/sqlx"
)

type IDeliveriesRepository interface {
	RepositoryGetAllDeliveries() ([]models.DeliveriesModel, error)
}

type DeliveriesRepository struct {
	*sqlx.DB
}

func InitializeRepoDeliveries(db *sqlx.DB) *DeliveriesRepository {
	dr := DeliveriesRepository{db}
	return &dr
}

func (r *DeliveriesRepository) RepositoryGetAllDeliveries() ([]models.DeliveriesModel, error) {
	result := []models.DeliveriesModel{}
	query := `SELECT * FROM deliveries`
	err := r.Select(&result, query)
	if err != nil {
		return nil, err
	}
	return result, nil
}