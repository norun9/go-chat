package main

import(
	"net/http"
	"html/template"
)

func index(w http.ResponseWriter, r *http.Request){
	files := []string{
		"templates/layout,html",
		"templates/navbar.html",
		"templates/index.html",
	}
	templates := template.Must(template.ParseFiles(files...))
	threads, err := data.Thread(); if err != nil {
		templates.ExecuteTemplate(w, "layouts", threads)
	}
}

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", index)
	mux.HnadleFunc("/err", err)
	mux.HnadleFunc("/login", login)
	mux.HnadleFunc("/logout", logout)
	mux.HnadleFunc("/signup", signup)
	mux.HnadleFunc("/signup_account", signup_account)
	mux.HnadleFunc("/authenticate", authenticate)

	mux.HnadleFunc("/thread/new", newTread)
	mux.HnadleFunc("/thread/create", createThread)
	mux.HnadleFunc("/thread/post", postThread)
	mux.HnadleFunc("/thread/read", readThread)

	server := http.Server{
		Addr: "127.0.0.1:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
