package routes

import (
	"example/bookstore/handlers"
	"example/bookstore/models"

	"github.com/gin-gonic/gin"
)

func RegisterBookRoutes(router *gin.Engine) {
	// Create an instance of DefaultBookStore
	store := &models.DefaultBookStore{}

	// Create a BookHandler using the DefaultBookStore
	handler := &handlers.BookHandler{Store: store}

	// Register routes with the handler methods
	router.GET("/books", handler.GetBooks)
	router.GET("/books/:id", handler.GetBooksById)
	router.POST("/books", handler.PostBooks)
}
