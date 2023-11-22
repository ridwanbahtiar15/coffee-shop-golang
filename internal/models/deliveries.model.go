package models

import "time"

type DeliveriesModel struct {
	Deliveries_id        string     `db:"deliveries_id"`
	Deliveries_name      string     `db:"deliveries_name" form:"deliveries_name" json:"deliveries_name"`
	Deliveries_cost      string     `db:"deliveries_cost" form:"deliveries_cost" json:"deliveries_cost"`
	Created_at *time.Time `db:"created_at"`
	Updated_at *time.Time `db:"updated_at"`
}