package models

import "time"

type PromosModel struct {
	Promos_id        string     `db:"promos_id"`
	Promos_name      string     `db:"promos_name" form:"promos_name" json:"promos_name"`
	Promos_start      string     `db:"promos_start" form:"promos_start" json:"promos_start"`
	Promos_end      string     `db:"promos_end" form:"promos_end" json:"promos_end"`
	Created_at *time.Time `db:"created_at"`
	Updated_at *time.Time `db:"updated_at"`
}