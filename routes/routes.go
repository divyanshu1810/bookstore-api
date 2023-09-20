package routes

import (
	"github.com/divyanshu1810/bookstore-api/routes/book_routes"
	"github.com/gorilla/mux"
)

func InitializeRoutes(router *mux.Router) {
	// Define your routes here
	book_routes.InitializeBookRoutes(router)
}
