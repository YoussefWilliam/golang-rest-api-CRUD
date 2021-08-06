package main

	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Book object struct
type Book struct {
	ID string `json:"id"`
	ISBN string `json:"isbn"`
	Title string `json:title"`
	Author *Author `json:"author"`
}

// Author object struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var books []Book

// Get all books
func getBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}
func getBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	// Get the id from the request params

	params:= mux.Vars(r) //map[id: 12]

	for _, book:= range books{
		if (book.ID == params["id"]){
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}
