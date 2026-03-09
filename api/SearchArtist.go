package api

import (
	"errors"
	"groupie-tracker/models"
)


func FetchArtistByName(Name string) (*models.Artist, error) {
	artists, err := FetchArtists()
	if err != nil {
		return nil, err
	}

	for i := range artists {
		if artists[i].Name == Name {
			return &artists[i], nil
		}
	}
	return nil, errors.New("artist not found")
}
