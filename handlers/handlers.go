package handlers

import (
	"html/template"
	"net/http"
	"strconv"

	"groupie-tracker/api"
	"groupie-tracker/models"
)

type PageData struct {
	Artists []models.Artist
	Dates   []models.Date
}

// HomeHandler returns an http.HandlerFunc that uses the provided artists
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	artists, err := api.FetchArtists()
	if err != nil {
		http.Error(w, "Artist error", 500)
		return
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


func ArtistHandler(w http.ResponseWriter, r *http.Request){
	idStr := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idStr)
	if err != nil{
		http.Error(w, "invalid artist ID", 400)
		return
	}

	artist, err := api.FetchArtistByID(id)
	if err != nil{
		http.Error(w, "artist not found", 404)
	}

	tmpl, err := template.ParseFiles()
	if err!= nil{
		http.Error(w, "template error", 500)
	}
	tmpl.Execute(w, artist)
}