package repositories

import (
	"coffee-shop-golang/internal/models"

	"github.com/jmoiron/sqlx"
)

type IRoleRepository interface {
	RepositoryGetAllRoles() ([]models.RolesModel, error)
}

type RolesRepository struct {
	*sqlx.DB
}

func InitializeRepoRoles(db *sqlx.DB) *RolesRepository {
	cr := RolesRepository{db}
	return &cr
}

func (r *RolesRepository) RepositoryGetAllRoles() ([]models.RolesModel, error) {
	result := []models.RolesModel{}
	query := `SELECT roles_id, roles_name FROM roles`
	err := r.Select(&result, query)
	if err != nil {
		return nil, err
	}
	return result, nil
}