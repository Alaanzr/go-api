package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var (
	err error
	DB *sql.DB
)

func init() {
	connStr := "user=postgres dbname=api password=postgres sslmode=disable"
	DB, err = sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}
}