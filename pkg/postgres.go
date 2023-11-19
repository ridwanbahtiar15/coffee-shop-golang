package pkg

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func PostgreSQLDB() (*sqlx.DB, error) {
	config := "host=localhost user=postgres password=12345 dbname=coffeeshop sslmode=disable"
	return sqlx.Connect("postgres", config)
}