package data

import(
	"database/sql"
	"log"
	"time"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "dbname=gochat sslmode=disable")
	if err != nil{
		log.Fatal(err)
	}
	return
}

type Tread struct {
	Id int
	Uuid string
	Topic string
	UserId string
	CreatedAt time.Time
}

