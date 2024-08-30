package handlers

import (
	"errors"
	"example/bookstore/models"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// MockBookStore is a mock implementation of the BookStore interface.
// It is used to simulate the behavior of the real BookStore in tests.
type MockBookStore struct {
	books []models.Book
}

// GetBooks returns the list of books in the mock store.
func (m *MockBookStore) GetBooks() ([]models.Book, error) {
	return m.books, nil
}

// GetBooksById searches for a book by ID in the mock store's list of books.
func (m *MockBookStore) GetBooksById(id string) (models.Book, error) {
	for _, book := range m.books {
		if book.ID == id {
			return book, nil
		}
	}
	return models.Book{}, errors.New("book not found")
}

// AddBook adds a new book to the mock store's list of books.
func (m *MockBookStore) AddBook(book *models.Book) error {
	m.books = append(m.books, *book)
	return nil
}

// TestGetBooks tests
func TestGetBooks(t *testing.T) {
	// Create a mock store with some books.
	mockStore := &MockBookStore{
		books: []models.Book{
			{ID: "1", Title: "Test Book 1", Author: "Author 1"},
			{ID: "2", Title: "Test Book 2", Author: "Author 2"},
		},
	}

	// Create a BookHandler with the mock store.
	handler := BookHandler{Store: mockStore}

	// Set up the HTTP recorder and Gin context for the test.
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)
	r.GET("/books", handler.GetBooks)
	req, _ := http.NewRequest("GET", "/books", nil)
	r.ServeHTTP(w, req)

	// Assert that the response status code is 200 OK.
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Test Book 1")
	assert.Contains(t, w.Body.String(), "Test Book 2")
}

// TestGetBooksById tests
func TestGetBooksById(t *testing.T) {
	// Create a mock store with one book.
	mockStore := &MockBookStore{
		books: []models.Book{
			{ID: "1", Title: "Test Book 1", Author: "Author 1"},
		},
	}
	handler := BookHandler{Store: mockStore}

	// Set up the HTTP recorder and Gin context for the test.
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)
	r.GET("/books/:id", handler.GetBooksById)
	req, _ := http.NewRequest("GET", "/books/1", nil)
	r.ServeHTTP(w, req)

	// Assert that the response status code is 200 OK and the body contains the expected book.
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Test Book 1")

	// Test for a book that does not exist.
	req, _ = http.NewRequest("GET", "/books/2", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert that the response status code is 404 Not Found for a non-existing book.
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "Book not found")
}

// TestPostBooks tests
func TestPostBooks(t *testing.T) {
	// Create an empty mock store.
	mockStore := &MockBookStore{}

	// Create a BookHandler with the mock store.
	handler := BookHandler{Store: mockStore}
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)
	r.POST("/books", handler.PostBooks)

	req, _ := http.NewRequest("POST", "/books", strings.NewReader(`{"title":"Test Book","author":"Test Author"}`))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	// Assert that the response status code is 201 Created.
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "Test Book")
	assert.Contains(t, w.Body.String(), "Test Author")
}
