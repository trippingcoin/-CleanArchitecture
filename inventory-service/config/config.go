package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func InitDB() *sql.DB {
	connStr := "host=localhost port=5433 user=postgres password=redmi dbname=inventory sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Cannot connect to DB:", err)
	}
	fmt.Println("Connected to DB.")
	return db
}
