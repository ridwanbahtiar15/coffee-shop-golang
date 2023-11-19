package models

import "time"

type CategoriesModel struct {
	Categories_id        string     `db:"categories_id"`
	Categories_name      string     `db:"categories_name" form:"categories_name" json:"categories_name"`
	Created_at *time.Time `db:"created_at"`
	Updated_at *time.Time `db:"updated_at"`
}