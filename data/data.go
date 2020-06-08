package data

import (
	"database/sql"
	"log"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "username=tomoyaueno dbname=gochat sslmode=disable")
	if err != nil{
		log.Fatal(err)
	}
	return
}
