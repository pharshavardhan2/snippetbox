package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

// application contains fields for dependency injection
// handlers are defined on application so they can use these dependencies
type application struct {
	infoLog  *log.Logger
	errorLog *log.Logger
}

func main() {
	// flag is used for describing and using command line arguments
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	// create shared loggers for application
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		infoLog:  infoLog,
		errorLog: errorLog,
	}

	// mux acts as a router mapping url patterns to their handlers
	mux := http.NewServeMux()

	// fileServer is used for serving static(files) over http
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	// http.Server is used for configuring server
	server := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Printf("Starting server on %s", *addr)
	// start the server
	err := server.ListenAndServe()
	errorLog.Fatal(err)
}
