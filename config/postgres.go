package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
)

func ConnectPostgres() *sql.DB {
	var DSN = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require TimeZone=Asia/Jakarta", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_DATABASE"), os.Getenv("DB_PORT"))

	db, err := sql.Open("postgres", DSN)
	if err != nil {
		panic(err)
	}

	return db
}
