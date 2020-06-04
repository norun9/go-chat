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

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string){
	var files []string
	for _, file := range filenames{
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}