package handlers

import (
	"html/template"
	"net/http"
	"strings"
	"strconv"

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

	//FILTERS
	minCreation,_ := strconv.Atoi(r.URL.Query().Get("minCreation"))
	maxCreation,_ := strconv.Atoi(r.URL.Query().Get("maxCreation"))

	//--album
	minAlbum,_ := strconv.Atoi(r.URL.Query().Get("minAlbum"))
	maxAlbum,_ := strconv.Atoi(r.URL.Query().Get("maxAlbum"))

	// --memebr & location --
	members := r.URL.Query()["members"]
	location := r.URL.Query()["Location"]


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
		//Filter  by Creation Date
		if minCreation != 0 && artist.CreationDate < minCreation {
			continue
		}
		if maxCreation != 0 && artist.CreationDate > maxCreation {
			continue
		}
		//Filter by first album year
		albumYearStr := artist.FirstAlbum[:4]
		albumYear, _ := strconv.Atoi(albumYearStr)

		if minAlbum != 0 && albumYear < minAlbum {
			continue
		}
		if maxAlbum != 0 && albumYear > maxAlbum {
			continue
		}
		//Filter by number of members
		if len(members) > 0 {
			match := false

			for _, m := range members {
				num,_ := strconv.Atoi(m)
				if len(artist.Members) == num {
					match = true
					break
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
		//Filter by location

		if len(location) > 0 {
			match := false

			for _, loc := range location {
				for key := range rel{
					if strings.Contains(strings.ToLower(key),strings.ToLower(loc)) {
						match = true
						break
					}
				}
			}
			if !match{
				continue
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
