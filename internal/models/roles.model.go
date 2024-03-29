package models

import "time"

type RolesModel struct {
	Roles_id        string     `db:"roles_id"`
	Roles_name      string     `db:"roles_name" form:"roles_name" json:"roles_name"`
	Created_at *time.Time `db:"created_at" json:"created_at,omitempty"`
	Updated_at *time.Time `db:"updated_at" json:"updated_at,omitempty"`
}