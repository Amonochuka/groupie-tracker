package handlers

import (
	"html/template"
	"net/http"
	"strings"

	// "strconv"
	// "strings"

	"groupie-tracker/api"
	"groupie-tracker/models"
)

type ArtistView struct {
	models.Artist
	Relations map[string][]string
}

// HomeHandler returns an http.HandlerFunc that uses the provided artists
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	query := strings.ToLower(r.URL.Query().Get("search"))
	artists, err := api.FetchArtists()
	if err != nil {
		http.Error(w, "Artist error", 500)
		return
	}
	relations, err := api.FetchRelations()
	if err != nil {
		http.Error(w, "RelationsError", 500)
	}
	// dates, err := api.GetDates()
	// if err != nil {
	// 	http.Error(w, "Dates error", 500)
	// 	return
	// }
	var data []ArtistView
	for _, artist := range artists {
		if query != "" {
			match := strings.Contains(strings.ToLower(artist.Name), query)
			for _, m := range artist.Members {
				if strings.Contains(strings.ToLower(m), query) {
					match = true
				}
			}
			if !match {
				continue
			}
		}
		var rel map[string][]string
		for _, r := range relations {
			if r.ID == artist.ID {
				rel = r.DatesLocations
				break
			}
		}

		data = append(data, ArtistView{
			Artist:    artist,
			Relations: rel,
		})
	}
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Template error", 500)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Execute error", 500)
	}
}
