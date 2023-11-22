package models

import "time"

type UsersModel struct {
	Users_id        		string     `db:"users_id"`
	Users_fullname      string     `db:"users_fullname" form:"users_fullname" json:"users_fullname"`
	Users_email      		string     `db:"users_email" form:"users_email" json:"users_email"`
	Users_password      string     `db:"users_password" form:"users_password" json:"users_password"`
	Users_phone 				string     `db:"users_phone" form:"users_phone" json:"users_phone"`
	Users_address      	string     `db:"users_address" form:"users_address" json:"users_address"`
	Users_image      		string     `db:"users_image" form:"users_image" json:"users_image"`
	Roles_id      			string     `db:"roles_id" form:"roles_id" json:"roles_id"`
	Roles_name      		string   		`db:"roles_name"`
	Created_at *time.Time `db:"created_at"`
	Updated_at *time.Time `db:"updated_at"`
	Deleted_at *time.Time `db:"deleted_at"`
	Is_active    			string     `db:"is_active"`
}