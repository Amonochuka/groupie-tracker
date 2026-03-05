package api

import (
	"encoding/json"
	"groupie-tracker/models"
	"net/http"
)

func FetchLocation() ([]models.Location, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data models.LocationResponse

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	return data.Index, nil

}

func fetchLocationByID(locations []models.Location, ID int) *models.Location {
	for _, location := range locations {
		if location.ID == ID {
			return &location
		}
	}
	return nil
}
