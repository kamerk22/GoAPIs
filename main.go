package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)




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
