package database

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

func NewDatabase() (*sql.DB, error) {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
