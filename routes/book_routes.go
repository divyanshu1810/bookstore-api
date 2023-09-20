package book_routes

import (
	"github.com/gorilla/mux"
	"github.com/yourusername/bookstore-api/handlers"
)

func InitializeBookRoutes(router *mux.Router) {
	// Define book-related routes here
	router.HandleFunc("/books", handlers.GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", handlers.GetBook).Methods("GET")
	router.HandleFunc("/books", handlers.CreateBook).Methods("POST")
	router.HandleFunc("/books/{id}", handlers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", handlers.DeleteBook).Methods("DELETE")
}
