package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	// create alias as _ as compiler complains as we don't use it directly
	// but init() in driver has to run to register with database/sql
	_ "github.com/go-sql-driver/mysql"
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
	dsn := flag.String("dsn", "web:web@/snippetbox?parseTime=true", "MYSQL data source name")
	flag.Parse()

	// create shared loggers for application
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

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
	err = server.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	// sql.Open() doesn't automatically open connection. Initializes db pool
	// check if everything works by opening a connection using Ping()
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
