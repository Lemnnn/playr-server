package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func InitDb() (*sql.DB, error) {
	godotenv.Load()
	connStr := os.Getenv("DB_URI")

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error initializing DB : %w", err)
	}

	return db, nil
}
