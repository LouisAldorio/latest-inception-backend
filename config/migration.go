package config

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

// migrate -database postgres://axiocxcetxtluk:10156336d1bf478f6ebd46655b0c0b80c576b2cf3941949cb1b6b8279f285fe9@ec2-3-211-245-154.compute-1.amazonaws.com:5432/d8vecr7uaocrsi?sslmode=require -path migration up

func MigrateSql() *migrate.Migrate {
	driver, err := postgres.WithInstance(ConnectPostgres(), &postgres.Config{})
	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migration",
		"postgres",
		driver,
	)
	if err != nil {
		panic(err)
	}

	log.Println("migration finished")
	return m
}
