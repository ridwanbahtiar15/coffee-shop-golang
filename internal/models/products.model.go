package models

import "time"

type ProductsModel struct {
	Products_id        string     `db:"products_id" valid:"-"`
	Products_name      string     `db:"products_name" form:"products_name" json:"products_name" valid:"required"`
	Products_price     string     `db:"products_price" form:"products_price" json:"products_price" valid:"required"`
	Products_desc      string     `db:"products_desc" form:"products_desc" json:"products_desc" valid:"required"`
	Products_stock     string     `db:"products_stock" form:"products_stock" json:"products_stock" valid:"required"`
	Categories_id     string     `db:"categories_id" form:"categories_id" json:"categories_id" valid:"required"`
	Products_image     string     `db:"products_image" json:"products_image" valid:"optional"`
	Created_at *time.Time `db:"created_at" valid:"-"`
	Updated_at *time.Time `db:"updated_at" valid:"-"`
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

type UpdateProductsModel struct {
	Products_name      string     `db:"products_name" form:"products_name" json:"products_name" valid:"optional"`
	Products_price     string     `db:"products_price" form:"products_price" json:"products_price" valid:"optional"`
	Products_desc      string     `db:"products_desc" form:"products_desc" json:"products_desc" valid:"optional"`
	Products_stock     string     `db:"products_stock" form:"products_stock" json:"products_stock" valid:"optional"`
	Categories_id     string     `db:"categories_id" form:"categories_id" json:"categories_id" valid:"optional"`
	Products_image     string     `db:"products_image" json:"products_image" valid:"optional"`
}