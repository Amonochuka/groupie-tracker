package api

import (
	"encoding/json"
	"errors"
	"groupie-tracker/models"
	"net/http"
)

var artists []models.Artist

// retrieve data from the Api and decode it to artists slice
func FetchArtists() ([]models.Artist, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&artists)
	return artists, err
}

func fetchArtistsByID(ID int) (models.Artist, error) {
	artists, err := FetchArtists()
	if err != nil {
		return models.Artist{}, err
	}

	for _, artist := range artists {
		if artist.ID == ID {
			return artist, nil
		}
	}
	return models.Artist{}, errors.New("artist not found")
}
