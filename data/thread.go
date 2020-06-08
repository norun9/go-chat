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

func Tread() (threads []Thread, err error){
	rows, err := Db.Query("SELECT count(*) FROM posts where thread_id = $1", thread.Id)
	if err != nil{
		return
	}
	for rows.Next(){
		th := new(Thread)
		if err = rows.Scan(th.Id, th.Uuid, th.Topic, th.UserId, th.CreateAt); err != nil{
			return
		}
		threads = append(threads, th)
	}
	rows.Close()
	return
}

func (thread *Thread) NumReplies() (count int) {
	rows, err := Db.Query("select count(*) from posts where thread_id = $1", thread.Id)
	if err != nil{
		return
	}
	for rows.Next(){
		if err = rows.Scan(&count); err != nil {
			return
		}

	}
	rows.Close()
	return
}