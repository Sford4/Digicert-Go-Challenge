package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/books", getAllBooks).Methods("GET")
	router.HandleFunc("/book", createSingleBook).Methods("POST")
	router.HandleFunc("/book/{id}", getSingleBook).Methods("GET")
	router.HandleFunc("/book/{id}", updateSingleBook).Methods("PUT")
	router.HandleFunc("/book/{id}", deleteSingleBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func main() {
	handleRequests()
}
