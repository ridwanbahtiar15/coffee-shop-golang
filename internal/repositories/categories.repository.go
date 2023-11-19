package repositories

import (
	"coffee-shop-golang/internal/models"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type CategoriesRepository struct {
	*sqlx.DB
}

func InitializeRepository(db *sqlx.DB) *CategoriesRepository {
	return &CategoriesRepository{db}
}

func (r *CategoriesRepository) RepsitoryGetAllCategories() ([]models.CategoriesModel, error) {
	result := []models.CategoriesModel{}
	query := `SELECT * FROM categories`
	err := r.Select(&result, query)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *CategoriesRepository) RepsitoryCreateCategories(body *models.CategoriesModel) (sql.Result, error) {
	query := `INSERT INTO categories (categories_name) VALUES (:categories_name)`
	result, err := r.NamedExec(query, body)
	if err != nil {
		return result, err 
	}
	return result, nil
}

func (r *CategoriesRepository) RepsitoryUpdateCategories(body *models.CategoriesModel, id string) (sql.Result, error) {
	query := `UPDATE categories SET categories_name=:categories_name WHERE categories_id =` + id
	result, err := r.NamedExec(query, body)
	if err != nil {
		return result, err 
	}
	return result, nil
}

func (r *CategoriesRepository) RepositoryDeleteCategories(id string) (sql.Result, error) {
	query := `DELETE FROM categories WHERE categories_id = $1`
	result, err := r.Exec(query, id)
	if err != nil {
		return result, err 
	}
	return result, nil
}