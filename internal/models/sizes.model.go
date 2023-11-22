package models

import "time"

type SizesModel struct {
	Sizes_id        string     `db:"sizes_id"`
	Sizes_name      string     `db:"sizes_name" form:"sizes_name" json:"sizes_name"`
	Created_at *time.Time `db:"created_at"`
	Updated_at *time.Time `db:"updated_at"`
}