package main

import (
	"log"
	"net/http"
	"go-webapp/handlers"
)

func main() {
	http.HandleFunc("/search/json", handlers.SearchInJSON)
	http.HandleFunc("/search/csv", handlers.SearchInCSV)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
