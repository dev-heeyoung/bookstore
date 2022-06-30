package main

import (
	"net/http"

	"github.com/dev-heeyoung/bookstore/pkg/models"
	"github.com/dev-heeyoung/bookstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	models.Init()
	r := mux.NewRouter()
	http.Handle("/", r)
	routes.RegisterBookStoreRoutes(r)
	http.ListenAndServe(":8000", r)

}
