package repositories

import (
	"coffee-shop-golang/internal/models"

	"github.com/jmoiron/sqlx"
)

type PromosRepository struct {
	*sqlx.DB
}

func InitializeRepoPromos(db *sqlx.DB) *PromosRepository {
	dr := PromosRepository{db}
	return &dr
}

func (r *PromosRepository) RepsitoryGetAllPromos() ([]models.PromosModel, error) {
	result := []models.PromosModel{}
	query := `SELECT * FROM promos`
	err := r.Select(&result, query)
	if err != nil {
		return nil, err
	}
	return result, nil
}