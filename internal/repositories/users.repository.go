package repositories

import (
	"coffee-shop-golang/internal/models"
	"fmt"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type IUsersRepository interface {
	RepositoryGetAllUsers(name string, page string, limit string, sort string) ([]models.UsersResponseModel, error)
	RepositoryUsersById(id string) ([]models.UsersResponseModel, error)
	RepositoryCreateUsers(body *models.UsersModel, hashedPassword string) (int, error)
	RepositoryUpdateUsers(body *models.UpdateUserModel, hashedPassword string, id string) (error)
	RepositoryUpdateImgUsers(usersImage string, id string) (error)
	RepositoryDeleteUsers(id string) (int, error)
	RepositoryCountUsers(name string) ([]string, error)
}

type UsersRepository struct {
	*sqlx.DB
}

func InitializeRepoUsers(db *sqlx.DB) *UsersRepository {
	cr := UsersRepository{db}
	return &cr
}

func (r *UsersRepository) RepositoryGetAllUsers(name string, page string, limit string, sort string) ([]models.UsersResponseModel, error) {
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

func (r *UsersRepository) RepositoryUsersById(id string) ([]models.UsersResponseModel, error) {
	result := []models.UsersResponseModel{}
	query := `SELECT u.users_fullname, u.users_email, u.users_phone, 
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

func (r *UsersRepository) RepositoryCreateUsers(body *models.UsersModel, hashedPassword string) (int, error) {
	var id int
	query := `INSERT INTO users (users_fullname, users_email, users_password, users_phone, users_address, users_image, roles_id) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING users_id`
	err := r.QueryRow(query, body.Users_fullname, body.Users_email, hashedPassword, body.Users_phone, body.Users_address, body.Users_image, body.Roles_id).Scan(&id)
	if err != nil {
		return 0, err 
	}
	return id, nil
}

func (r *UsersRepository) RepositoryUpdateUsers(body *models.UpdateUserModel, hashedPassword string, id string) (error) {
	query := `UPDATE users SET users_fullname = $1, users_password = $2, users_phone = $3, users_address = $4, users_image = $5, updated_at = NOW() WHERE users_id = $6`
	values := []any{body.Users_fullname, hashedPassword, body.Users_phone, body.Users_address, body.Users_image, id}
	_, err := r.Exec(query, values...)
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

func (r *UsersRepository) RepositoryDeleteUsers(id string) (int, error) {
	var res int = 0
	query := `UPDATE users SET deleted_at = NOW() WHERE users_id = $1`
	result, err := r.Exec(query, id)
	if err != nil {
		return 0, err 
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		res = 0
		return res, nil
	}
	res = 1
	return res, nil
}

// func (r *UsersRepository) RepositoryGetFilterUsers(name string, page string, limit string, sort string) ([]models.UsersResponseModel, error) {
// 	newPage, _ := strconv.Atoi("1")
// 	newLimit, _ := strconv.Atoi("99")

// 	if page != "" {
// 		newPage, _ = strconv.Atoi(page) 
// 	}
// 	if limit != "" {
// 		newLimit, _ = strconv.Atoi(limit) 
// 	}

// 	result := []models.UsersResponseModel{}
// 	query := `SELECT u.users_id, u.users_fullname, u.users_email, u.users_phone, 
// 						u.users_address, u.users_image, r.roles_name
// 						FROM users u
// 						JOIN roles r ON u.roles_id = r.roles_id`

// 	if name != "" {
// 		query += ` WHERE u.users_fullname LIKE $1`
// 		switch sort {
// 		case "asc":
// 			query += ` ORDER BY u.users_fullname ASC LIMIT $2 OFFSET $3`
// 		case "desc":
// 			query += ` ORDER BY u.users_fullname DESC LIMIT $2 OFFSET $3`
// 		default:
// 			query += ` ORDER BY u.users_fullname ASC LIMIT $2 OFFSET $3`
// 		}
// 		offset := newPage * newLimit - newLimit;
// 		err := r.Select(&result, query, fmt.Sprintf("%%%s%%", name), newLimit, strconv.Itoa(offset))
// 		if err != nil {
// 			return nil, err
// 		}
// 		return result, nil
// 	}

// 	switch sort {
// 	case "asc":
// 		query += ` ORDER BY u.users_fullname ASC LIMIT $1 OFFSET $2`
// 	case "desc":
// 		query += ` ORDER BY u.users_fullname DESC LIMIT $1 OFFSET $2`
// 	default:
// 		query += ` ORDER BY u.users_fullname ASC LIMIT $1 OFFSET $2`
// 	}
// 	offset := newPage * newLimit - newLimit;
// 	err := r.Select(&result, query, newLimit, strconv.Itoa(offset))
// 	if err != nil {
// 		return nil, err
// 	}
// 	return result, nil
// }

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