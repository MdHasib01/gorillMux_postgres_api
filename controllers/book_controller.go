package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/MdHasib01/gorillMux_postgres_api/db"
	"github.com/MdHasib01/gorillMux_postgres_api/models"
	"github.com/MdHasib01/gorillMux_postgres_api/repository"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
    books, err := repository.GetAllBooks(db.DB)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(books)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
    var book models.Book
    if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    if err := repository.CreateBook(db.DB, &book); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(book)
}
