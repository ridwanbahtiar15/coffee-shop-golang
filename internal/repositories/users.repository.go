package repositories

import (
	"coffee-shop-golang/internal/models"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type UsersRepository struct {
	*sqlx.DB
}

func InitializeRepoUsers(db *sqlx.DB) *UsersRepository {
	cr := UsersRepository{db}
	return &cr
}

func (r *UsersRepository) RepositoryGetAllUsers() ([]models.UsersResponseModel, error) {
	result := []models.UsersResponseModel{}
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

func (r *UsersRepository) RepositoryUsersById(id string) ([]models.UsersGetByIdResponseModel, error) {
	result := []models.UsersGetByIdResponseModel{}
	query := `SELECT u.users_fullname, u.users_email, u.users_password, u.users_phone, 
						u.users_address, u.users_image, r.roles_name 
						FROM users u
						JOIN roles r on u.roles_id = r.roles_id
						WHERE users_id = $1`
	err := r.Select(&result, query, id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *UsersRepository) RepositoryCreateUsers(body *models.UsersModel) (error) {
	query := `INSERT INTO users (users_fullname, users_email, users_password, users_phone, users_address, users_image, roles_id) VALUES (:users_fullname, :users_email, :users_password, :users_phone, :users_address, :users_image, :roles_id)`
	_, err := r.NamedExec(query, body)
	if err != nil {
		return err 
	}
	return nil
}

func (r *UsersRepository) RepositoryUpdateUsers(body *models.UsersModel, id string) (error) {
	query := `UPDATE users SET users_fullname=:users_fullname, users_password=:users_password, users_phone=:users_phone, users_address=:users_address, users_image=:users_image, updated_at=NOW() WHERE users_id =` + id
	_, err := r.NamedExec(query, body)
	if err != nil {
		return err 
	}
	return nil
}

func (r *UsersRepository) RepositoryUpdateImgUsers(usersImage string, id string) (error) {
	query := `UPDATE users SET users_image = $1 WHERE users_id = $2`
	_, err := r.Exec(query, usersImage,  id)
	if err != nil {
		return err 
	}
	return nil
}

func (r *UsersRepository) RepositoryDeleteUsers(id string) (sql.Result, error) {
	query := `UPDATE users SET deleted_at = NOW() WHERE users_id = $1`
	result, err := r.Exec(query, id)
	if err != nil {
		return nil, err 
	}
	return result, nil
}

func (r *UsersRepository) RepositoryGetFilterUsers(name string, page string, limit string, sort string) ([]models.UsersResponseModel, error) {
	newPage, _ := strconv.Atoi("1")
	newLimit, _ := strconv.Atoi("99")

	if page != "" {
		newPage, _ = strconv.Atoi(page) 
	}
	if limit != "" {
		newLimit, _ = strconv.Atoi(limit) 
	}

	result := []models.UsersResponseModel{}
	query := `SELECT u.users_id, u.users_fullname, u.users_email, u.users_phone, 
						u.users_address, u.users_image, r.roles_name
						FROM users u
						JOIN roles r ON u.roles_id = r.roles_id`

	if name != "" {
		query += ` WHERE u.users_fullname LIKE $1`
		switch sort {
		case "asc":
			query += ` ORDER BY u.users_fullname ASC LIMIT $2 OFFSET $3`
		case "desc":
			query += ` ORDER BY u.users_fullname DESC LIMIT $2 OFFSET $3`
		default:
			query += ` ORDER BY u.users_fullname ASC LIMIT $2 OFFSET $3`
		}
		offset := newPage * newLimit - newLimit;
		err := r.Select(&result, query, fmt.Sprintf("%%%s%%", name), newLimit, strconv.Itoa(offset))
		if err != nil {
			return nil, err
		}
		return result, nil
	}

	switch sort {
	case "asc":
		query += ` ORDER BY u.users_fullname ASC LIMIT $1 OFFSET $2`
	case "desc":
		query += ` ORDER BY u.users_fullname DESC LIMIT $1 OFFSET $2`
	default:
		query += ` ORDER BY u.users_fullname ASC LIMIT $1 OFFSET $2`
	}
	offset := newPage * newLimit - newLimit;
	err := r.Select(&result, query, newLimit, strconv.Itoa(offset))
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *UsersRepository) RepositoryCountUsers(name string) ([]string, error) {
	count := []string{}
	query := `SELECT COUNT(*) FROM users u`

	if name != "" {
		query += ` WHERE u.users_fullname LIKE $1`
		err := r.Select(&count, query, fmt.Sprintf("%%%s%%", name))
		if err != nil {
			return nil, err
		}
		return count, nil
	}

	err := r.Select(&count, query)
		if err != nil {
			return nil, err
		}
		return count, nil
}