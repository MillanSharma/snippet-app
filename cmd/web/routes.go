package main

import (
	"net/http"
)

func (app *snippetBox) routes() *http.ServeMux {

	mux := http.NewServeMux()

	// file server
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// routes handler
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	return mux
}
