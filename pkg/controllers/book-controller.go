package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/talha-yazar/Bookstore-Management-API/pkg/models"
	"github.com/talha-yazar/Bookstore-Management-API/pkg/utils"
)

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()

	res, _ := json.Marshal(newBooks)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	bookId := params["bookId"]

	ID, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	book, err := models.GetBookById(ID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	res, _ := json.Marshal(book)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}

	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["bookId"]
	bookId, err := strconv.ParseInt(id, 0, 0)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	UpdateBook := &models.Book{ID: bookId}

	utils.ParseBody(r, UpdateBook)
	isUpdateSuccessful := UpdateBook.UpdateBookById()

	if !isUpdateSuccessful {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		fmt.Println("Book not found")
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeleteBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["bookId"]

	ID, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	isDeleteSuccessful := models.DeleteBookById(ID)

	if !isDeleteSuccessful {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		fmt.Println("Book not found")
	}

	w.WriteHeader(http.StatusNoContent)
}
