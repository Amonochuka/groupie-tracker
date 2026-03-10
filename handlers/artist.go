package handlers

import (
	"html/template"
	"net/http"
	"strconv"

	// "strings"

	"groupie-tracker/api"
)



func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid artist ID", 400)
		return
	}

	artist, err := api.FetchArtistByID(id)
	if err != nil {
		http.Error(w, "Artists error", 500)
		return
	}

	relations, err := api.FetchRelations()
	if err != nil {
		http.Error(w, "Relations error", 500)
		return
	}
	var rel map[string][]string
	for _, r := range relations {
		if r.ID == id {
			rel = r.DatesLocations
			break
		}
	}
	data := struct {
		ID           int
		Name         string
		Image        string
		Members      []string
		CreationDate int
		FirstAlbum   string
		Relations    map[string][]string
	}{
		ID:           artist.ID,
		Name:         artist.Name,
		Image:        artist.Image,
		Members:      artist.Members,
		CreationDate: artist.CreationDate,
		FirstAlbum:   artist.FirstAlbum,
		Relations:    rel,
	}

	tmpl, err := template.ParseFiles("templates/artist.html")
	if err != nil {
		http.Error(w, "template error", 500)
	}
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), 500)
	}
}
