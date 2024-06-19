package app

import (
	"net/http"
)

func (a *App) RegisterRoutes() http.Handler {

	mux := http.NewServeMux()
	mux.HandleFunc("/", a.IndexHandler)
	// mux.HandleFunc("/home", a.HomeHandler)

	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/*", http.StripPrefix("/static/", fs))

	return mux
}
