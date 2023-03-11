package main

import (
	"golang-backend-dev/crud-api-with-db/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.BookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8081", r))
}
