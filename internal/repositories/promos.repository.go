package repositories

import (
	"coffee-shop-golang/internal/models"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type IPromosRepository interface {
	RepsitoryGetAllPromos(page string, limit string) ([]models.PromosModel, error)
	RepositoryGetPromosById(id string) ([]models.PromosModel, error)
	RepsitoryCreatePromos(body *models.PromosModel) (error)
	RepsitoryUpdatePromos(body *models.UpdatePromosModel, id string) (error)
	RepositoryDeletePromos(id string) (int64, error)
	RepositoryCountPromos() ([]string, error)
}

type PromosRepository struct {
	*sqlx.DB
}

func InitializeRepoPromos(db *sqlx.DB) *PromosRepository {
	dr := PromosRepository{db}
	return &dr
}

func (r *PromosRepository) RepsitoryGetAllPromos(page string, limit string) ([]models.PromosModel, error) {
	// result := []models.PromosModel{}
	// query := `SELECT * FROM promos`
	// err := r.Select(&result, query)
	// if err != nil {
	// 	return nil, err
	// }
	// return result, nil

	newPage, _ := strconv.Atoi("1")
	newLimit, _ := strconv.Atoi("99")

	if page != "" {
		newPage, _ = strconv.Atoi(page) 
	}
	if limit != "" {
		newLimit, _ = strconv.Atoi(limit) 
	}

	result := []models.PromosModel{}
	query := `SELECT * FROM promos WHERE promos_id != '0' LIMIT $1 OFFSET $2`
	offset := newPage * newLimit - newLimit;
		err := r.Select(&result, query, newLimit, strconv.Itoa(offset))
		if err != nil {
			return nil, err
		}
		return result, nil
}

func (r *PromosRepository) RepositoryGetPromosById(id string) ([]models.PromosModel, error) {
	result := []models.PromosModel{}
	query := `SELECT * FROM promos WHERE promos_id = $1`
	err := r.Select(&result, query, id)
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

func (r *PromosRepository) RepsitoryUpdatePromos(body *models.UpdatePromosModel, id string) (error) {
	query := `UPDATE promos SET promos_name=:promos_name, promos_start=:promos_start, promos_end=:promos_end, updated_at = NOW() WHERE promos_id = ` + id;
	_, err := r.NamedExec(query, body)
	if err != nil {
		return err
	}
	return nil
}

func (r *PromosRepository) RepositoryDeletePromos(id string) (int64, error) {
	var res int64 = 1
	query := `DELETE FROM promos WHERE promos_id = $1`
	result, err := r.Exec(query, id)
	if err != nil {
		return 0, err 
	}
	
	if rows, _ := result.RowsAffected(); rows == 0 {
		res = 0
		return res, nil
	}
	return res, nil
}

func (r *PromosRepository) RepositoryGetFilterPromos(page string, limit string) ([]models.PromosModel, error) {
	newPage, _ := strconv.Atoi("1")
	newLimit, _ := strconv.Atoi("99")

	if page != "" {
		newPage, _ = strconv.Atoi(page) 
	}
	if limit != "" {
		newLimit, _ = strconv.Atoi(limit) 
	}

	result := []models.PromosModel{}
	query := `SELECT * FROM promos WHERE promos_id != '0' LIMIT $1 OFFSET $2`
	offset := newPage * newLimit - newLimit;
		err := r.Select(&result, query, newLimit, strconv.Itoa(offset))
		if err != nil {
			return nil, err
		}
		return result, nil
}

func (r *PromosRepository) RepositoryCountPromos() ([]string, error) {
	count := []string{}
	query := `SELECT COUNT(*) FROM promos WHERE promos_id != '0'`
	err := r.Select(&count, query)
		if err != nil {
			return nil, err
		}
		return count, nil
}