package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/yourusername/bookstore-api/routes"
)

func main() {
	router := mux.NewRouter()
	routes.InitializeRoutes(router)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", router)
}
