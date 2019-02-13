package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// Book struct ( Model )
type Books struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author struct ( Model )
type Author struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func main() {
	// Init router
	r := mux.NewRouter()

	// Route handlers
	r.HandleFunc("/api/books/", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBooks).Methods("GET")
	r.HandleFunc("/api/books/", createBooks).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBooks).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBooks).Methods("DELETE")

	// Start the server
	log.Fatal(http.ListenAndServe(":8000", r))
}
