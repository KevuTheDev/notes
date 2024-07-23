package main

import (
	"log"
	"net/http"
)

func main() {
	// Setting up server
	mux := http.NewServeMux()

	// routes
	mux.HandleFunc("/", home)
	mux.HandleFunc("/note/view", noteView)
	mux.HandleFunc("/note/create", noteCreate)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
