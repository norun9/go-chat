package main

import (
	"encoding/json"
	"errors"
	"fmt"
	// 	"github.com/sausheong/gwp/Chapter_2_Go_ChitChat/chitchat/data"
	"github.com/mushahiroyuki/gowebprog/ch02/chitchat/data"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

func session(w http.ResponseWriter, r *http.Request)(sess data.Session, err errors){
	cookie, err := r.Cookie("_coolie")
	if err != nil{
		sess = data.Session{Uuid: cookie.Value}
		if ok, _ := sess.Check(); !ok {
			err = errors.New("Invalid session")
		}
	}
	return
}