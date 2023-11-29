package models

import "time"

type OrdersModel struct {
	Orders_id        					string     `db:"orders_id" form:"orders_id" json:"orders_id"`
	Users_id      						string     `db:"users_id" form:"users_id" json:"users_id"`
	Deliveries_id     				string     `db:"deliveries_id" form:"deliveries_id" json:"deliveries_id"`
	Promos_id      						string     `db:"promos_id" form:"promos_id" json:"promos_id"`
	Payment_methods_id    		string     `db:"payment_methods_id" form:"payment_methods_id" json:"payment_methods_id"`
	Orders_status      				string     `db:"orders_status" form:"orders_status" json:"orders_status"`
	Orders_total      				string     `db:"orders_total" form:"orders_total" json:"orders_total"`
	Orders_products_id    		string     `db:"orders_products_id" form:"orders_products_id" json:"orders_products_id"`
	Products_id        				string     `db:"products_id" form:"products_id" json:"products_id"`
	Sizes_id        					string     `db:"sizes_id" form:"sizes_id" json:"sizes_id"`
	Orders_products_qty   		string     `db:"orders_products_qty" form:"orders_products_qty" json:"orders_products_qty"`
	Orders_products_subtotal  string     `db:"orders_products_subtotal" form:"orders_products_subtotal" json:"orders_products_subtotal"`
	Hot_or_ice  string     `db:"hot_or_ice" form:"hot_or_ice" json:"hot_or_ice"`
	Created_at *time.Time `db:"created_at"`
	Updated_at *time.Time `db:"updated_at"`
}


type OrdersResponseModel struct {
	Orders_id        					string     `db:"orders_id"`
	Users_id      						string     		`db:"users_id"`
	Deliveries_id     				string     `db:"deliveries_id"`
	Promos_name      						string     `db:"promos_name"`
	Payment_methods_id    		string     `db:"payment_methods_id"`
	Orders_status      				string     `db:"orders_status"`
	Orders_total      				string     `db:"orders_total"`
}