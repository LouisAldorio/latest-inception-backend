package config

import (
	"database/sql"
	"fmt"
	"os"

	"myapp/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

//ConnectDB connect GO to database
func ConnectDB() (*gorm.DB, *sql.DB) {
	var dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))

	var mysql = mysql.New(mysql.Config{
		DSN: dsn,
	})

	db, err := gorm.Open(mysql, &gorm.Config{
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
