package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

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
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)

    //find book sent in list of books by comparing ID, return whole book on finding
    for _, item := range books {
        if item.ID == params["id"]{
            fmt.Println(item)
            json.NewEncoder(w).Encode(item)
            return
        }
    }

    //if no book found return empty book JSON
    json.NewEncoder(w).Encode(&Book{})
}

func createBook(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json")
    var book Book

    //data sent is in request body, bust be decoded (parsed)
    _ = json.NewDecoder(r.Body).Decode(&book)
    book.ID = strconv.Itoa(rand.Intn(10000000))
    books = append(books, book)
    json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)

    for index, item := range books {
        if item.ID == params["id"]{
            books = append(books[:index], books[index+1:]...)
            var book Book

            //data sent is in request body, bust be decoded (parsed)
            _ = json.NewDecoder(r.Body).Decode(&book)
            book.ID = params["id"]
            books = append(books, book)
            json.NewEncoder(w).Encode(book)
            return
        }
    }
    json.NewEncoder(w).Encode(books)
}

func deleteBook(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)

    for index, item := range books {
        if item.ID == params["id"]{
            books = append(books[:index], books[index+1:]...)
            break
        }
    }
    json.NewEncoder(w).Encode(books)
}


func main() {
   
    //init router
    router := mux.NewRouter()

    //Mock data, Author is another struct inside Book struct
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