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

func (r *PromosRepository) RepsitoryCreatePromos(body *models.PromosModel) (error) {
	query := `INSERT INTO promos (promos_name, promos_start, promos_end) values (:promos_name, :promos_start, :promos_end)`;
	_, err := r.NamedExec(query, body)
	if err != nil {
		return err
	}
	return nil
}

func (r *PromosRepository) RepsitoryUpdatePromos(body *models.PromosModel, id string) (error) {
	query := `UPDATE promos SET promos_name=:promos_name, promos_start=:promos_start, promos_end=:promos_end, updated_at = NOW() WHERE promos_id = $4`;
	_, err := r.NamedExec(query, body)
	if err != nil {
		return err
	}
	return nil
}

func (r *PromosRepository) RepositoryDeletePromos(id string) (error) {
	query := `DELETE FROM promos WHERE promos_id = $1`
	_, err := r.Exec(query, id)
	if err != nil {
		return err 
	}
	return nil
}