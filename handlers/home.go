package handlers

import (
	"html/template"
	"net/http"
	// "strconv"
	// "strings"

	"groupie-tracker/api"
	"groupie-tracker/models"
)

type PageData struct {
	Artists []models.Artist
	Dates   []models.Date
}

// HomeHandler returns an http.HandlerFunc that uses the provided artists
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// query:=strings.ToLower(r.URL.Query().Get("search"))
	artists, err := api.FetchArtists()
	if err != nil {
		http.Error(w, "Artist error", 500)
		return
	}
// relations ,err:= api.FetchRelations()
if err != nil{
	http.Error(w,"RelationsError",500)
}
	dates, err := api.GetDates()
	if err != nil {
		http.Error(w, "Dates error", 500)
		return
	}
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Template error", 500)
		return
	}
	data := PageData{
		Artists: artists,
		Dates:   dates,
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Execute error", 500)
	}
}


