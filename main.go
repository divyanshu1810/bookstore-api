package main

import (
	"fmt"
	"net/http"

	"github.com/divyanshu1810/bookstore-api/routes"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	routes.InitializeRoutes(router)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", router)
}
