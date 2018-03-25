package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "bledar"
	password = "bledi1"
	dbname   = "nba"
)

var conn string

func Connection () *sql.DB{

	conn = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, _ = sql.Open("postgres", conn)
	_, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err)
	}
	return db
}
