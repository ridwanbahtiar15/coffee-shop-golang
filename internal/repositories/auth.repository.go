package repositories

import (
	"coffee-shop-golang/internal/models"

	"github.com/jmoiron/sqlx"
)

type RepositoryAuth struct {
	*sqlx.DB
}

func InitializeRepoAuth(db *sqlx.DB) *RepositoryAuth {
	cr := RepositoryAuth{db}
	return &cr
}

func (r *RepositoryAuth) RepositoryRegisterUsers(body *models.GetUserInfoModel, hashedPassword string) (error) {
	query := `INSERT INTO users (users_fullname, users_email, users_password) VALUES ($1, $2, $3)`
	values := []any{body.Users_fullname, body.Users_email, hashedPassword}
	_, err := r.Exec(query, values...)
	if err != nil {
		return err 
	}
	return nil
}

func (r *RepositoryAuth) GetUsers(body *models.GetUserInfoModel) ([]models.GetUserInfoModel, error) {
	query := "SELECT users_id, users_fullname, users_email, users_password, roles_id FROM users WHERE users_email = $1"
	values := []any{body.Users_email}
	result := []models.GetUserInfoModel{}
	if err := r.Select(&result, query, values...); err != nil {
		return nil, err
	}
	return result, nil
}

func (r *RepositoryAuth) InsertJwt(usersId string, tokenJwt string) (error) {
	query := `INSERT INTO users_tokenjwt (users_id, token_jwt) VALUES ($1, $2)`
	_, err := r.Exec(query, usersId, tokenJwt)
	if err != nil {
		return err 
	}
	return nil
}

func (r *RepositoryAuth) DeleteJwt(token string) (error) {
	query := `DELETE from users_tokenjwt WHERE token_jwt = $1`
	_, err := r.Exec(query, token)
	if err != nil {
		return err 
	}
	return nil
}

// func (r *RepositoryAuth) GetUsersJwt(tokenJwt string) ([]models.JwtUsers, error) {
// 	result := []models.JwtUsers{}
// 	query := `SELECT users_id, token_jwt FROM users_tokenjwt WHERE token_jwt = $1`
// 	if err := r.Select(&result, query, tokenJwt); err != nil {
// 		return nil, err
// 	}
// 	return result, nil
// }