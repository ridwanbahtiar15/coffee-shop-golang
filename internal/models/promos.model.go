package models

import "time"

type PromosModel struct {
	Promos_id        string     `db:"promos_id" valid:"-"`
	Promos_name      string     `db:"promos_name" form:"promos_name" json:"promos_name" valid:"required"`
	Promos_start      string     `db:"promos_start" form:"promos_start" json:"promos_start" valid:"required"`
	Promos_end      string     `db:"promos_end" form:"promos_end" json:"promos_end" valid:"required"`
	Created_at *time.Time `db:"created_at" valid:"-"`
	Updated_at *time.Time `db:"updated_at" valid:"-"`
}

type UpdatePromosModel struct {
	Promos_name      string     `db:"promos_name" form:"promos_name" json:"promos_name" valid:"optional"`
	Promos_start      string     `db:"promos_start" form:"promos_start" json:"promos_start" valid:"optional"`
	Promos_end      string     `db:"promos_end" form:"promos_end" json:"promos_end" valid:"optional"`
}