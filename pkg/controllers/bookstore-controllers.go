package controllers

import (
	"encoding/json"
	"fmt"
	"golang-backend-dev/crud-api-with-db/pkg/models"
	"golang-backend-dev/crud-api-with-db/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	response, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	bookId := mux.Vars(r)["bookId"]
	bookIdInt, err := strconv.ParseInt(bookId, 10, 64)
	if err != nil {
		fmt.Println("error while parsing bookId")
	}
	book, _ := models.GetBookById(bookIdInt)
	response, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	createdBook := models.Book{}
	utils.ParseBody(r, &createdBook)
	book := createdBook.CreateBook()
	response, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	bookIdStr := mux.Vars(r)["bookId"]
	bookIdInt, err := strconv.ParseInt(bookIdStr, 10, 64)
	if err != nil {
		fmt.Println("error while parsing bookId")
	}
	book := models.DeleteBook(bookIdInt)
	response, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	updatedBook := models.Book{}
	utils.ParseBody(r, &updatedBook)

	bookIdStr := mux.Vars(r)["bookId"]
	bookIdInt, err := strconv.ParseInt(bookIdStr, 10, 64)
	if err != nil {
		fmt.Println("error while parsing bookId")
	}
	bookDetails, db := models.GetBookById(bookIdInt)

	if updatedBook.Name != "" {
		bookDetails.Name = updatedBook.Name
	}
	if updatedBook.Author != "" {
		bookDetails.Author = updatedBook.Author
	}
	if updatedBook.Publication != "" {
		bookDetails.Publication = updatedBook.Publication
	}
	db.Save(bookDetails)

	response, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
