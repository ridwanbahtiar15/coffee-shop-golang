package models

import "time"

type PaymentmethodsModel struct {
	Payment_methods_id        string     `db:"payment_methods_id"`
	Payment_methods_name      string     `db:"payment_methods_name" form:"payment_methods_name" json:"payment_methods_name"`
	Created_at *time.Time `db:"created_at"`
	Updated_at *time.Time `db:"updated_at"`
}