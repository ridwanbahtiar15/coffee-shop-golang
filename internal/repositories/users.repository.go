package repositories

import (
	"coffee-shop-golang/internal/models"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type UsersRepository struct {
	*sqlx.DB
}

func InitializeRepoUsers(db *sqlx.DB) *UsersRepository {
	cr := UsersRepository{db}
	return &cr
}

func (r *UsersRepository) RepsitoryGetAllUsers() ([]models.UsersModel, error) {
	result := []models.UsersModel{}
	query := `SELECT u.users_id, u.users_fullname, u.users_email, u.users_phone, 
						u.users_address, u.users_image, r.roles_name 
						FROM users u
						JOIN roles r on u.roles_id = r.roles_id
						WHERE deleted_at is NULL
						ORDER BY u.users_id ASC`
	err := r.Select(&result, query)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *UsersRepository) RepsitoryUsersById(id string) ([]models.UsersModel, error) {
	result := []models.UsersModel{}
	query := `SELECT u.users_fullname, u.users_email, u.users_password, u.users_phone, 
						u.users_address, u.users_image, u.roles_id FROM users u WHERE users_id = $1`
	err := r.Select(&result, query, id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *UsersRepository) RepsitoryCreateUsers(body *models.UsersModel) (sql.Result, error) {
	query := `INSERT INTO users (users_fullname, users_email, users_password, users_phone, users_address, users_image, roles_id) VALUES (:users_fullname, :users_email, :users_password, :users_phone, :users_address, :users_image, :roles_id)`
	result, err := r.NamedExec(query, body)
	if err != nil {
		return result, err 
	}
	return result, nil
}

func (r *UsersRepository) RepsitoryUpdateUsers(body *models.UsersModel, id string) (sql.Result, error) {
	query := `UPDATE users SET users_fullname=:users_fullname, users_password=:users_password, users_phone=:users_phone, users_address=:users_address, users_image=:users_image, updated_at=NOW() WHERE users_id =` + id
	result, err := r.NamedExec(query, body)
	if err != nil {
		return result, err 
	}
	return result, nil
}

func (r *UsersRepository) RepositoryDeleteUsers(id string) (sql.Result, error) {
	query := `UPDATE users SET deleted_at = NOW() WHERE users_id = $1`
	result, err := r.Exec(query, id)
	if err != nil {
		return result, err 
	}
	return result, nil
}