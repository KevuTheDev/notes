package main

import (
	"io/fs"
	"path/filepath"
	"text/template"
	"time"

	"github.com/KevuTheDev/notes/internal/models"
	"github.com/KevuTheDev/notes/ui"
)

func humanReadableDate(t time.Time) string {
	if t.IsZero() {
		return ""
	}

	return t.UTC().Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanReadableDate,
}

type templateData struct {
	CurrentYear int
	Note        *models.Note
	Notes       []*models.Note
	Form        any
	Flash       string
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := fs.Glob(ui.Files, "html/pages/*.templ")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		patterns := []string{
			"html/base.templ",
			"html/partials*.templ",
			page,
		}

		ts, err := template.New(name).Funcs(functions).ParseFS(ui.Files, patterns...)
		if err != nil {
			return nil, err
		}
		cache[name] = ts
	}

	return cache, nil
}
