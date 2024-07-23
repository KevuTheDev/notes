package main

import (
	"fmt"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func (app *application) noteView(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	fmt.Fprintf(w, "Display a specific snippet with ID %s...", id)
}

func (app *application) noteCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating a note"))
}
