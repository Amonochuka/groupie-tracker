package api

import (
	"encoding/json"
	"errors"
	"groupie-tracker/models"
	"net/http"
)

func FetchLocations(client HTTPClient) ([]models.Location, error) {
	req, _ := http.NewRequest("GET", "https://groupietrackers.herokuapp.com/api/locations", nil)
	resp, err := client.Do(req)
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

func fetchLocationByID(client HTTPClient, ID int) (*models.Location, error) {
	locations, err := FetchLocations(client)
	if err != nil {
		return nil, err
	}
	for i := range locations {
		if locations[i].ID == ID {
			return &locations[i], nil
		}
	}
	return nil, errors.New("location not found")
}
