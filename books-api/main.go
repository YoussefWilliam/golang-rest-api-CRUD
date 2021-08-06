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
// Create a book

func createBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)


	book.ID = strconv.Itoa(rand.Intn(100))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)	

}
// Update a book
func updateBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index,item := range books {
		if(item.ID == params["id"]){
			books = append(books[:index],books[index+1:]...)
			break
		}
	}
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)


	book.ID = params["id"]
	books = append(books, book)
	json.NewEncoder(w).Encode(book)	

}
// Delete a book
func deleteBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index,item := range books {
		if(item.ID == params["id"]){
			books = append(books[:index],books[index+1:]...)
			break
		}
	}
}
