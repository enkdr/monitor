package app

import (
	"fmt"
	"net/http"
	
)

type pageData struct {
	Title string
}

func (a *App) IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("index handler")
	pageData := pageData{
		Title: "Home",
	}

	err := a.templates.ExecuteTemplate(w, "index.html", pageData)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// func (a *App) HomeHandler(w http.ResponseWriter, r *http.Request) {

// 	fmt.Println("home handler")
// 	pageData := pageData{
// 		Title: "Home",
// 	}

// 	err := s.templates.ExecuteTemplate(w, "index.html", pageData)

// 	if err != nil {
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}

// }
