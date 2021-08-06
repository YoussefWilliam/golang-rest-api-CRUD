package main

import (
	"fmt"
	"log"
	"net/http"
)


func homePage(write http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(write, "Welcome to the homepage")
	fmt.Println("Endpoint is::: Homepage")
}

func handleRequest(){
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func main() {
	handleRequest()
}