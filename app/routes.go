package app

import (
	"net/http"
)

func (a *App) RegisterRoutes() http.Handler {

	mux := http.NewServeMux()
	mux.HandleFunc("/", a.IndexHandler)

	fs := http.FileServer(http.Dir("app/static"))
	mux.Handle("/static/*", http.StripPrefix("/static/", fs))

	return mux
}
