package api

import (
	"encoding/json"
	"errors"
	"groupie-tracker/models"
	"net/http"
)


// retrieve data from the Api and decode it to artists slice
func FetchArtists() ([]models.Artist, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	
	var artists []models.Artist

	err = json.NewDecoder(resp.Body).Decode(&artists)
	if err != nil{
		return nil,err
	}
	return artists, nil
}

func FetchArtistByID(ID int) (*models.Artist, error) {
	artists, err := FetchArtists()
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
