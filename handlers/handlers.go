package handlers

import (
	"html/template"
	"net/http"

	"groupie-tracker/models"
)

// HomeHandler returns an http.HandlerFunc that uses the provided artists
func HomeHandler(artists []models.Artist) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			http.Error(w, "Template error", 500)
			return
		}

		err = tmpl.Execute(w, artists)
		if err != nil {
			http.Error(w, "Execute error", 500)
		}
	}
}
