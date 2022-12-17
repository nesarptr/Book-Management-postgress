package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nesarptr/Book-Management-postgress/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":3000", r))
}
