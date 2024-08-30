package models

import (
	"database/sql"
	"example/bookstore/database"
	"net/http"
)

// Book represents a book entity with basic information.
type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

// BookStore is an interface that defines the methods
type BookStore interface {
	GetBooks() ([]Book, error)
	GetBooksById(id string) (Book, error)
	AddBook(book *Book) error
}

// DefaultBookStore is a concrete implementation of the BookStore interface.
type DefaultBookStore struct{}

// BookError is a custom error
type BookError struct {
	Code    int
	Message string
}

// Error implements the error interface for BookError.
func (e *BookError) Error() string {
	return e.Message
}

// NewBookError creates a new BookError with a given status code and message.
func NewBookError(code int, message string) *BookError {
	return &BookError{
		Code:    code,
		Message: message,
	}
}

// runQuery is a reusable function that runs an exec query on the database (e.g., INSERT, UPDATE, DELETE).
// It takes a query string and arguments for the query.
func runQuery(query string, args ...interface{}) (sql.Result, error) {
	result, err := database.DB.Exec(query, args...)
	return result, err
}

// runQueryRows is a reusable function that runs a query on the database and processes multiple rows.
// It takes a query string, a function to scan each row, and arguments for the query.
func runQueryRows(query string, scanFunc func(*sql.Rows) error, args ...interface{}) error {
	rows, err := database.DB.Query(query, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		if err := scanFunc(rows); err != nil {
			return err
		}
	}
	return rows.Err()
}

func (db *DefaultBookStore) GetBooks() ([]Book, error) {
	var books []Book

	// Use the runQueryRows function to execute the query and scan the results.
	err := runQueryRows("SELECT * FROM books", func(rows *sql.Rows) error {
		var book Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price); err != nil {
			return err
		}
		books = append(books, book)
		return nil
	})

	if err != nil {
		return nil, err
	}
	return books, nil
}

func (db *DefaultBookStore) GetBooksById(id string) (Book, error) {
	var book Book
	query := "SELECT * FROM books WHERE id = ?"
	row := database.DB.QueryRow(query, id)
	// Scan the result into the book struct.
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Price)
	if err == sql.ErrNoRows {
		// Return a custom error with a 404 status code if no rows are found.
		return book, NewBookError(http.StatusNotFound, "Book not found")
	}
	if err != nil {

		return book, NewBookError(http.StatusInternalServerError, err.Error())
	}
	return book, nil
}

func (db *DefaultBookStore) AddBook(book *Book) error {
	// Prepare the query to insert a new book into the database.
	query := "INSERT INTO books (title, author, price) VALUES (?,?,?)"
	_, err := runQuery(query, book.Title, book.Author, book.Price)
	return err
}
