package data

import (
	"time"
)

type Thread struct {
	Id        int
	Uuid      string
	Topic     string
	UserId    int
	CreatedAt time.Time
}

func Tread() (threads []Tread, err error){
	rows, err := Db.Query("SELECT count(*) FROM posts where thread_id = $1", thread.Id)
	if err != nil{
		return
	}
	for rows.Next(){
		th := new(Tread)
		if err = rows.Scan(th.Id, th.Uuid, th.Topic, th.UserId, th.CreateAt); err != nil{
			return
		}
		threads = append(threads, th)
	}
	rows.Close()
	return
}