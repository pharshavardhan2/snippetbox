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

	// http.Server is used for configuring server
	server := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", *addr)
	// start the server
	err := server.ListenAndServe()
	errorLog.Fatal(err)
}
