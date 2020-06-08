package main

import(
	"net/http"
	"github.com/mushahiroyuki/gowebprog/ch02/chitchat/data"
)

//func index(w http.ResponseWriter, r *http.Request){
//	threads, err := data.Thread(); if err == nil {
//		_, err := session(w, r)
//		public_tmpl_files := []string{
//			"templates/layout,html",
//			"templates/public.navbar.html",
//			"templates/index.html",
//		}
//		private_tmpl_files := []string{
//			"templates/layout,html",
//			"templates/private.navbar.html",
//			"templates/index.html",
//		}
//		var templates *template.Template
//		if err != nil {
//			templates = template.Must(template.ParseFiles(private_tmpl_files...))
//		} else {
//			templates = template.Must(template.ParseFiles(public_tmpl_files...))
//		}
//		templates.ExecuteTemplate(w, "layouts", threads)
//	}
//}

func index(w http.ResponseWriter, r *http.Request){
	threads, err := data.Threads(); if err == nil {
		_, err := session(w, r)
		if err != nil {
			generateHTML(w, threads, "layout", "public.navbar", "index")
		}else{
			generateHTML(w, threads, "layout", "private.navbar", "index")
		}
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