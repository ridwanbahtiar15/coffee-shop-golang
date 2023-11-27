package models

import "time"

type ProductsModel struct {
	Products_id        string     `db:"products_id"`
	Products_name      string     `db:"products_name" form:"products_name" json:"products_name"`
	Products_price     string     `db:"products_price" form:"products_price" json:"products_price"`
	Products_desc      string     `db:"products_desc" form:"products_desc" json:"products_desc"`
	Products_stock     string     `db:"products_stock" form:"products_stock" json:"products_stock"`
	Products_image     string     `db:"products_image" form:"products_image" json:"products_image"`
	Categories_id      string     `db:"categories_id" form:"categories_id" json:"categories_id"`
	Created_at *time.Time `db:"created_at"`
	Updated_at *time.Time `db:"updated_at"`
}

type ProductsResponseModel struct {
	Products_id        string     `db:"products_id"`
	Products_name      string     `db:"products_name"`
	Products_price     string     `db:"products_price"`
	Products_desc      string     `db:"products_desc"`
	Products_stock     string     `db:"products_stock"`
	Products_image     string     `db:"products_image"`
	Categories_id      string     `db:"categories_id"`
	Categories_name 	 string 		`db:"categories_name"`
}

type Meta struct {
	Page int
	TotalData int
	Next string
	Prev string
}
