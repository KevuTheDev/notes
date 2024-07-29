package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/KevuTheDev/notes/internal/models"
)

type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	notes         models.NoteModelInterface
	templateCache map[string]*template.Template
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")

	dsn := flag.String("dsn", "web:OBNRGPh9QRRD!bwY@Tx&WLUD@/notebook?parseTime=true", "MySQL data source name")

	flag.Parse()

	// Logging information
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close()

	templateCache, err := newTemplateCache()
	if err != nil {
		errorLog.Fatal(err)
	}

	app := &application{
		infoLog:       infoLog,
		errorLog:      errorLog,
		notes:         &models.NoteModel{DB: db},
		templateCache: templateCache,
	}
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
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/note/view", app.noteView)
	mux.HandleFunc("/note/create", app.noteCreate)

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
