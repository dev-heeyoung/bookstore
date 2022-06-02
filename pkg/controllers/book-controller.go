package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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

func GetBook(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()
	res, _ := json.Marshal(books)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["bookId"]
	parsedId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("err while parsing")
	}
	book, _ := models.GetBookById(parsedId)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["bookId"]
	parsedId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("err while parsing")
	}
	book := models.DeleteBook(parsedId)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["bookId"]
	parsedId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("err while parsing")
	}
	book, db := models.GetBookById(parsedId)

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
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
