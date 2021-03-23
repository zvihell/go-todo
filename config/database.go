package config

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var DBClient *sql.DB

func InitDB() {
	db, err := sql.Open("postgres", "user=postgres password=88888888 dbname=gotodo sslmode=disable")

	if err != nil {
		panic(err)
	}
	DBClient = db
}
