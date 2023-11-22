package pkg

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func PostgreSQLDB() (*sqlx.DB, error) {
	err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_SSLMODE"))

	return sqlx.Connect("postgres", config)
}