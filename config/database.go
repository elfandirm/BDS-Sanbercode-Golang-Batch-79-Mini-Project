package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func ConnectDatabase() (*sql.DB, error) {
	dsn := os.Getenv("DATABASE_URL")

	if dsn == "" {
		dsn = "postgres://postgres:1234567890@localhost:5432/bioskop_db"
	}

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("gagal ping database: %w", err)
	}

	return db, nil
}