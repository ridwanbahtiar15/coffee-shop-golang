package repositories

import (
	"coffee-shop-golang/internal/models"

	"github.com/jmoiron/sqlx"
)

type ICategoriesRepository interface {
	RepositoryGetAllCategories() ([]models.CategoriesModel, error)
}

type CategoriesRepository struct {
	*sqlx.DB
}

func InitializeRepoCategories(db *sqlx.DB) *CategoriesRepository {
	// return &CategoriesRepository{db}
	cr := CategoriesRepository{db}
	return &cr
}

// func (cr CategoriesRepository) InitializeRepository(db *sqlx.DB) *CategoriesRepository {
// 	return &CategoriesRepository{db}
// }

func (r *CategoriesRepository) RepositoryGetAllCategories() ([]models.CategoriesModel, error) {
	result := []models.CategoriesModel{}
	query := `SELECT * FROM categories`
	err := r.Select(&result, query)
	if err != nil {
		return nil, err
	}
	return result, nil
}
