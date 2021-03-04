package config

import (
	"database/sql"
	"fmt"
	"os"

	"myapp/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DSN = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require TimeZone=Asia/Jakarta", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_DATABASE"), os.Getenv("DB_PORT"))

//ConnectDB connect GO to database
func ConnectDB() (*gorm.DB, *sql.DB) {
	var postgres = postgres.New(postgres.Config{
		DSN:                  DSN,
		PreferSimpleProtocol: true,
	})

	db, err := gorm.Open(postgres, &gorm.Config{
		Logger: logger.InitLog(),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		panic(err)
	}

	sql, _ := db.DB()

	return db, sql
}
