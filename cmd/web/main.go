package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

type application struct {
	debug bool
}

func main() {
	fmt.Println("Hello World")

	addr := flag.String("addr", ":4000", "HTTP network address")

	flag.Parse()

	app := &application{
		debug: false,
	}

	srv := &http.Server{
		Addr:         *addr,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Starting server on :4000")
	err := srv.ListenAndServe()
	log.Fatal(err)
}
