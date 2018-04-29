package main

import (
	"net/http"
	pages "portfolio/handlers"

	"github.com/gorilla/mux"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello, world!"))
}

func main() {
	r := mux.NewRouter()

	// API responses
	r.HandleFunc("/hello", sayHello)

	// Webpages
	r.HandleFunc("/todo", pages.TodoPageHandler)
	r.HandleFunc("/contact", pages.ContactFormPageHandler)

	// Static pages
	s := http.FileServer(http.Dir("assets/"))
	r.PathPrefix("/").Handler(s)
	http.Handle("/", r)

	http.ListenAndServe(":8080", r)
}
