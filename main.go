package main

import (
	"log"
	"myapp/config"

	"github.com/joho/godotenv"
)

func main2() {
	godotenv.Load()

	m := config.MigrateSql()
	defer m.Close()

	err := m.Up()
	if err != nil {
		v, d, _ := m.Version()
		if d {
			m.Force(int(v) - 1)
		}
		panic(err)
	}

	log.Println("Migration finished")
}
