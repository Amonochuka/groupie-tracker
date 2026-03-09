package handlers

import (
	"html/template"
	"net/http"
	"strconv"
	// "strings"

	"groupie-tracker/api"
	// "groupie-tracker/models"
)

// type PageData struct {
// 	Artists []models.Artist
// 	Dates   []models.Date
// }

func ArtistHandler(w http.ResponseWriter, r *http.Request){
	idStr := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idStr)
	if err != nil{
		http.Error(w, "invalid artist ID", 400)
		return
	}

	artist, err := api.FetchArtistByID(id)
	// relations,_:=api.FetchRelations()
	
	if err != nil{
		http.Error(w, "artist not found", 404)
	}

	tmpl, err := template.ParseFiles("templates/artist.html")
	if err!= nil{
		http.Error(w, "template error", 500)
	}
	tmpl.Execute(w, artist)
}