package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/KevuTheDev/notes/internal/models"
	"github.com/julienschmidt/httprouter"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	notes, err := app.notes.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := app.newTemplateData(r)
	data.Notes = notes

	app.render(w, http.StatusOK, "home.templ", data)
}

func (app *application) noteView(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	// We can then use the ByName() method to get the value of the "id" named
	// parameter from the slice and validate it as normal.
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	note, err := app.notes.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	data := app.newTemplateData(r)
	data.Note = note
	app.render(w, http.StatusOK, "view.templ", data)
}

func (app *application) noteCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", http.MethodPost) // let client know it allows for POST requests at this end point
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)

		return
	}

	w.Write([]byte("Creating new note..."))
}
