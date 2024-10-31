package repository

import (
	"database/sql"

	"github.com/MdHasib01/gorillMux_postgres_api/models"
)

func GetAllBooks(db *sql.DB) ([]models.Book, error) {
    rows, err := db.Query("SELECT id, title, author, description FROM books")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var books []models.Book
    for rows.Next() {
        var book models.Book
        if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Description); err != nil {
            return nil, err
        }
        books = append(books, book)
    }
    return books, nil
}

func CreateBook(db *sql.DB, book *models.Book) error {
    query := "INSERT INTO books (title, author, description) VALUES ($1, $2, $3) RETURNING id"
    return db.QueryRow(query, book.Title, book.Author, book.Description).Scan(&book.ID)
}
