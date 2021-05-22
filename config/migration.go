package config

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func MigrateSql() *migrate.Migrate {
	driver, err := mysql.WithInstance(connectSql(), &mysql.Config{})
	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migration",
		"mysql",
		driver,
	)
	if err != nil {
		panic(err)
	}

	return m
}

func connectSql() *sql.DB {
	var dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	return db
}
