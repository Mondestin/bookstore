package handlers

import (
	"example/bookstore/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// BookHandler uses a BookStore interface for its operations.
type BookHandler struct {
	Store models.BookStore
}

// GetBooks handles the GET request to retrieve all books.
func (h *BookHandler) GetBooks(c *gin.Context) {
	books, err := h.Store.GetBooks()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, books)
}

// GetBooksById handles the GET request to retrieve a single book by its ID.
func (h *BookHandler) GetBooksById(c *gin.Context) {
	// Get the book ID from the request parameters.
	id := c.Param("id")
	book, err := h.Store.GetBooksById(id)
	if err != nil {
		if err.Error() == "book not found" {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		} else {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}

// PostBooks handles the POST request to add a new book.
func (h *BookHandler) PostBooks(c *gin.Context) {
	var newBook models.Book
	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	if err := h.Store.AddBook(&newBook); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, newBook)
}
