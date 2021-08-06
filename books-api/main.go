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
