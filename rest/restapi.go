package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Book struct --> Model
type Book struct {
    ID string `json:"id"`
    Isbn string `json:"isbn"`
    Title string `json:"title"`
    Author *Author `json:"author"`
}

// Init books var as a slice Book struct
var books []Book

//Author struct --> Model
type Author struct {
    Firstname string `json:"firstname"`
    Lastname string `json:"lastname"`
}

func getBooks(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request){

}

func createBook(w http.ResponseWriter, r *http.Request){

}

func updateBook(w http.ResponseWriter, r *http.Request){

}

func deleteBook(w http.ResponseWriter, r *http.Request){

}


func main() {
   
    //init router
    router := mux.NewRouter()

    //Mock data
    books = append(books, Book{ID: "1", Isbn: "438227", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "454555", Title: "Book Two", Author: &Author{Firstname: "Steve", Lastname: "Smith"}})


    //route handlers/Endpoints
    router.HandleFunc("/books", getBooks).Methods("GET")
    router.HandleFunc("/book/{id}", getBook).Methods("GET")
    router.HandleFunc("/book", createBook).Methods("POST")
    router.HandleFunc("/book/{id}", updateBook).Methods("PUT")
    router.HandleFunc("/book/{id}", deleteBook).Methods("DELETE")

    //Listen to server, end with error on failure
    log.Fatal(http.ListenAndServe(":3000", router))

}