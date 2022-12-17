package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nesarptr/Book-Management-postgress/pkg/models"
	"github.com/nesarptr/Book-Management-postgress/pkg/utils"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	createBook := &models.Book{}
	utils.ParseBody(r, createBook)
	Book := createBook.CreateBook()
	res, _ := json.MarshalIndent(Book, "", "\t")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.MarshalIndent(newBooks, "", "\t")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	ID, _ := strconv.ParseUint(mux.Vars(r)["bookId"], 0, 0)
	Book, _ := models.GetBookById(uint(ID))
	res, _ := json.MarshalIndent(Book, "", "\t")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	ID, _ := strconv.ParseUint(mux.Vars(r)["bookId"], 0, 0)
	Book := models.DeleteById(uint(ID))
	res, _ := json.MarshalIndent(Book, "", "\t")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	updatedBook := &models.Book{}
	utils.ParseBody(r, updatedBook)
	ID, _ := strconv.ParseUint(mux.Vars(r)["bookId"], 0, 0)
	Book, db := models.GetBookById(uint(ID))

	if updatedBook.Name != "" {
		Book.Name = updatedBook.Name
	}

	if updatedBook.Author != "" {
		Book.Author = updatedBook.Author
	}

	if updatedBook.Publication != "" {
		Book.Publication = updatedBook.Publication
	}

	db.Save(Book)

	res, _ := json.MarshalIndent(Book, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
