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

/* Book Structures [Model] */
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Init books var as a slice Book struct
var books []Book

/*
*	Every route handler must intake a response and request
*	w http.ResponseWriter
*	r *http.Request
 */

// Getting all the books
func getBooks(w http.ResponseWriter, r *http.Request) {
	// Ensuring that the data comes out in JSON format
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(books)
}

// Retrieving Single Book by id
func getBook(w http.ResponseWriter, r *http.Request) {
	// Ensuring that the data comes out in JSON format
	w.Header().Set("Content-Type", "application/json")

	// Get any paramaters
	param := mux.Vars(r)

	// Loop through books JSON and find by id
	for _, item := range books {
		if item.ID == param["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	// Should return a single Book(structname one)
	json.NewEncoder(w).Encode(&Book{})
}

// Creating a new book record
func createBooks(w http.ResponseWriter, r *http.Request) {
	// Ensuring that the data comes out in JSON format
	w.Header().Set("Content-Type", "application/json")

	// var <variableName> <struct name>
	var book Book

	// Setting the data from the request into the struct for our use
	_ = json.NewDecoder(r.Body).Decode(&book)

	// Mock ID - Not safe for Production
	book.ID = strconv.Itoa(rand.Intn(10000000))
	books = append(books, book)

	// Should return a single Book(structname one)
	json.NewEncoder(w).Encode(book)
}

// Updating a book record
func updateBooks(w http.ResponseWriter, r *http.Request) {
	// Ensuring that the data comes out in JSON format
	w.Header().Set("Content-Type", "application/json")

	// Get the id
	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)

			// var <variableName> <struct name>
			var book Book

			// Setting the data from the request into the struct for our use
			_ = json.NewDecoder(r.Body).Decode(&book)

			// Mock ID - Not safe for Production
			book.ID = params["id"]
			books = append(books, book)

			// Should return a single Book(structname one)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}

// Delete book
func deleteBooks(w http.ResponseWriter, r *http.Request) {
	// Ensuring that the data comes out in JSON format
	w.Header().Set("Content-Type", "application/json")

	// Get the id
	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	fmt.Println("Starting REST API")
	// Initialise Router
	r := mux.NewRouter()

	// Sample Data
	books = append(books,
		Book{
			ID:    "1",
			Isbn:  "21892831781",
			Title: "Book One",
			Author: &Author{
				Firstname: "John",
				Lastname:  "Smith"}})

	books = append(books,
		Book{
			ID:    "2",
			Isbn:  "57892831781",
			Title: "Book Two",
			Author: &Author{
				Firstname: "Steven",
				Lastname:  "Jagesh"}})

	// Route Handlers - Establish our endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBooks).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBooks).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBooks).Methods("DELETE")

	// Running the Server
	// Wrapping in log.Fatal so that we get a log of what the error is incase of failure
	log.Fatal(http.ListenAndServe(":8000", r))

}
