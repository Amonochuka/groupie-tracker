package api

import (
	"encoding/json"
	"errors"
	"groupie-tracker/models"
	"net/http"
)

var DefaultClient = &http.Client{}

// retrieve data from the Api and decode it to artists slice
func FetchArtists(client HTTPClient) ([]models.Artist, error) {
	req, _ := http.NewRequest("GET", "https://groupietrackers.herokuapp.com/api/artists", nil)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var artists []models.Artist

	err = json.NewDecoder(resp.Body).Decode(&artists)
	if err != nil {
		return nil, err
	}
	return artists, nil
}

func FetchArtistByID(client HTTPClient, ID int) (*models.Artist, error) {
	artists, err := FetchArtists(client)
	if err != nil {
		return nil, err
	}

	for i := range artists {
		if artists[i].ID == ID {
			return &artists[i], nil
		}
	}
	return nil, errors.New("artist not found")
}
