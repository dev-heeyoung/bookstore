package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/dev-heeyoung/bookstore/pkg/cache"
	"github.com/dev-heeyoung/bookstore/pkg/models"
	"github.com/dev-heeyoung/bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	newBook := models.Book{}
	utils.ParseBody(r, &newBook)
	book := newBook.CreateBook()

	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()
	res, _ := json.Marshal(books)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["bookId"]

	book := cache.GetBook(id)
	if book.Id != "" {
		book, _ = models.GetBookById(id)
		if book.Id != "" {
			w.WriteHeader(http.StatusNotFound)
			return
		} else {
			cache.SaveBook(book)
		}
	}
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["bookId"]
	cache.DeleteBook(id)
	book := models.DeleteBook(id)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["bookId"]

	cache.DeleteBook(id)

	book, db := models.GetBookById(id)

	var updateBook = models.Book{}
	utils.ParseBody(r, &updateBook)

	if book.Author != "" {
		book.Author = updateBook.Author
	}
	if book.Name != "" {
		book.Name = updateBook.Name
	}
	if book.Publication != "" {
		book.Publication = updateBook.Publication
	}
	db.Save(book)
	cache.SaveBook(book)

	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
