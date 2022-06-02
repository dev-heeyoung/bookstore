package main

import (
	"log"
	"net/http"

	"github.com/dev-heeyoung/bookstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	http.Handle("/", r)
	routes.RegisterBookStoreRoutes(r)
	log.Fatal(http.ListenAndServe(":8000", r))
}
