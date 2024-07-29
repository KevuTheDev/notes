package main

import (
	"net/http"

	"github.com/KevuTheDev/notes/ui"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.notFound(w)
	})

	fileServer := http.FileServer(http.FS(ui.Files))

	router.Handler(http.MethodGet, "/static/*filepath", fileServer)

	dynamic := alice.New()

	router.Handler(http.MethodGet, "/", dynamic.ThenFunc(app.home))
	router.Handler(http.MethodGet, "/note/view/:id", dynamic.ThenFunc(app.noteView))
	router.Handler(http.MethodGet, "/note/create", dynamic.ThenFunc(app.noteCreate))

	return dynamic.Then(router)
}
