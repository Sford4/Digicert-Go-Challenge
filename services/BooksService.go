package services

import (
	"LibraryAPI/data"
	database "LibraryAPI/database"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var db database.Database

func InitDB() error {
	envErr := godotenv.Load()
	if envErr != nil {
		fmt.Printf("Error loading credentials: %v", envErr)
	}

	var (
		password     = os.Getenv("MSSQL_DB_PASSWORD")
		user         = os.Getenv("MSSQL_DB_USER")
		port         = os.Getenv("MSSQL_DB_PORT")
		databaseName = os.Getenv("MSSQL_DB_DATABASE")
	)

	connectionString := fmt.Sprintf("user id=%s;password=%s;port=%s;database=%s", user, password, port, databaseName)

	sqlObj, connectionError := sql.Open("mssql", connectionString)
	if connectionError != nil {
		fmt.Println(fmt.Errorf("error opening database: %v", connectionError))
	}

	db = database.Database{
		SqlDb: sqlObj,
	}

	return nil
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := db.RetrieveBooks()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("Failed to retrieve books"))
		return
	}
	w.WriteHeader(http.StatusOK)
	stringBody, _ := json.Marshal(books)
	_, _ = w.Write([]byte(stringBody))

	return
}

func GetSingleBook(w http.ResponseWriter, r *http.Request) {
	// get path params
	vars := mux.Vars(r)
	bookID, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("Missing book id"))
		return
	}

	book, err := db.RetrieveBook(bookID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte("Failed to retrieve book"))
		return
	}
	w.WriteHeader(http.StatusOK)
	stringBody, _ := json.Marshal(book)
	_, _ = w.Write([]byte(stringBody))
	return
}

func CreateSingleBook(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	// Unmarshal string into structs.
	var newBook data.Book
	err := decoder.Decode(&newBook)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Failed to create book"))
		return
	}
	id, err := db.CreateBook(&newBook)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte("Failed to create book"))
		return
	}
	w.WriteHeader(http.StatusOK)
	stringID := strconv.Itoa(id)
	_, _ = w.Write([]byte(stringID))
	return
}

func UpdateSingleBook(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	// Unmarshal string into structs.
	var newBook data.Book
	err := decoder.Decode(&newBook)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Failed to create book"))
		return
	}
	err = db.UpdateBook(&newBook)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Failed to update book"))
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}

func DeleteSingleBook(w http.ResponseWriter, r *http.Request) {
	// get path params
	vars := mux.Vars(r)
	bookID, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("Missing book id"))
		return
	}

	err := db.DeleteBook(bookID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Failed to delete book"))
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}
