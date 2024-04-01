package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	// mux acts as a router mapping url patterns to their handlers
	mux := http.NewServeMux()

	// fileServer is used for serving static(files) over http
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	return mux
}
