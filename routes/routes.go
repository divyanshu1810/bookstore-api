package routes

import (
	"github.com/gorilla/mux"
	"github.com/yourusername/bookstore-api/routes/book_routes"
)

func InitializeRoutes(router *mux.Router) {
	// Define your routes here
	book_routes.InitializeBookRoutes(router)
}
