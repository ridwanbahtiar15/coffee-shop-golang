package models

import "time"

type UsersModel struct {
	Users_id        		string     `db:"users_id" valid:"-"`
	Users_fullname      string     `db:"users_fullname" form:"users_fullname" json:"users_fullname" valid:"required"`
	Users_email      		string     `db:"users_email" form:"users_email" json:"users_email" valid:"email,required"`
	Users_password      string     `db:"users_password" form:"users_password" json:"users_password" valid:"required,length(3|12)"`
	Users_phone 				string     `db:"users_phone" form:"users_phone" json:"users_phone" valid:"optional"`
	Users_address      	string     `db:"users_address" form:"users_address" json:"users_address" valid:"optional"`
	Users_image      		string     `db:"users_image" json:"users_image" valid:"optional"`
	Roles_id      			string     `db:"roles_id" form:"roles_id" json:"roles_id" valid:"required"`
	Roles_name      		string   		`db:"roles_name" valid:"-"`
	Created_at *time.Time `db:"created_at" valid:"-"`
	Updated_at *time.Time `db:"updated_at" valid:"-"`
	Deleted_at *time.Time `db:"deleted_at" valid:"-"`
	Is_active    			string     `db:"is_active" valid:"-"`
}

type UsersResponseModel struct {
	Users_id        		string     `db:"users_id" json:"users_id,omitempty"`
	Users_fullname      string     `db:"users_fullname"`
	Users_email      		string     `db:"users_email"`
	Users_password      string     `db:"users_password" json:"users_password,omitempty"`
	Users_phone 				string     `db:"users_phone"`
	Users_address      	string     `db:"users_address"`
	Users_image      		string     `db:"users_image"`
	Roles_name      		string   		`db:"roles_name"`
}

// type UsersGetByIdResponseModel struct {
// 	UsersResponseModel
	
// }

type GetUserInfoModel struct {
	Users_id        		string     `db:"users_id" valid:"-"`
	Users_fullname      string     `db:"users_fullname" form:"users_fullname" json:"users_fullname" valid:"-"`
	Users_email      		string     `db:"users_email" form:"users_email" json:"users_email" valid:"email,required"`
	Users_password      string     `db:"users_password" form:"users_password" json:"users_password" valid:"required"`
	Roles_id      			string     `db:"roles_id" valid:"-"`
}

type JwtUsers struct {
	Users_id string `db:"users_id" form:"users_id" json:"users_id"`
	Token_jwt string `db:"token_jwt" form:"token_jwt" json:"token_jwt"`
}

type UpdateUserModel struct {
	Users_fullname      string     `db:"users_fullname" form:"users_fullname" json:"users_fullname" valid:"optional"`
	Users_email      		string     `db:"users_email" form:"users_email" json:"users_email" valid:"email,optional"`
	Users_password      string     `db:"users_password" form:"users_password" json:"users_password" valid:"optional,length(3|12)"`
	Users_phone 				string     `db:"users_phone" form:"users_phone" json:"users_phone" valid:"optional"`
	Users_address      	string     `db:"users_address" form:"users_address" json:"users_address" valid:"optional"`
	Users_image      		string     `db:"users_image" json:"users_image" valid:"optional"`
}