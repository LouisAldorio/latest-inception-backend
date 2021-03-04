package config

import (
	"database/sql"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
)

func ConnectPostgres() *sql.DB {
	db, err := sql.Open("postgres", DSN)
	if err != nil {
		panic(err)
	}

	return db
}
