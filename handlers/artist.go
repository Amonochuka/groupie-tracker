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
		RenderError(w, http.StatusBadRequest, "Invalid Artist ID.")
		return
	}

	artist, err := api.FetchArtistByID(http.DefaultClient, id)
	if err != nil {
		RenderError(w, http.StatusNotFound, "Artist not found")
		return
	}

	relations, err := api.FetchRelations(http.DefaultClient)
	if err != nil {
		RenderError(w, http.StatusInternalServerError, "Failed to load concert information")
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
		RenderError(w,
			http.StatusInternalServerError,
			"Internal server error.",
		)
	}
	if err := tmpl.Execute(w, data); err != nil {
		RenderError(w,
			http.StatusInternalServerError,
			"Failed to render artist page.",
		)
	}
}
