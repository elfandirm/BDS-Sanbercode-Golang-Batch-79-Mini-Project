package config

import (
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func ConnectDatabase() (*sql.DB, error) {
	dsn := "postgres://postgres:1234567890@localhost:5432/bioskop_db"

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}