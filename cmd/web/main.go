package main

import (
	"log"
	"net/http"
)

func main() {
	// Setting up server
	mux := http.NewServeMux()

	// Set up a file server that delivers files from the "./ui/static" directory.
	// Ensure that the path specified for the http.Dir function is relative
	// to the project's root directory.
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Use the mux.Handle() function to register the file server as the handler
	// for all URL paths that start with "/static/". For matching paths, we
	// strip the "/static" prefix before the request reaches the file server.
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// routes
	mux.HandleFunc("/", home)
	mux.HandleFunc("/note/view", noteView)
	mux.HandleFunc("/note/create", noteCreate)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
