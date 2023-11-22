package repositories

import (
	"coffee-shop-golang/internal/models"

	"github.com/jmoiron/sqlx"
)

type RolesRepository struct {
	*sqlx.DB
}

func InitializeRepoRoles(db *sqlx.DB) *RolesRepository {
	cr := RolesRepository{db}
	return &cr
}

func (r *RolesRepository) RepsitoryGetAllRoles() ([]models.RolesModel, error) {
	result := []models.RolesModel{}
	query := `SELECT * FROM roles`
	err := r.Select(&result, query)
	if err != nil {
		return nil, err
	}
	return result, nil
}