package main

import (
	"LibraryAPI/services"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequests() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/books", services.GetAllBooks).Methods("GET")
	router.HandleFunc("/book", services.CreateSingleBook).Methods("POST")
	router.HandleFunc("/book/{id}", services.GetSingleBook).Methods("GET")
	router.HandleFunc("/book/{id}", services.UpdateSingleBook).Methods("PUT")
	router.HandleFunc("/book/{id}", services.DeleteSingleBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func main() {
	services.InitDB()
	handleRequests()
}
