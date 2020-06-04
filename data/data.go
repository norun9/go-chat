package data

import(
	"time"
)

type Tread struct {
	Id int
	Uuid string
	Topic string
	UserId string
	CreatedAt time.Time
}

