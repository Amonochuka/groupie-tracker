package api

import (
	"encoding/json"
	"errors"
	"groupie-tracker/models"
	"net/http"
)

func FetchLocations() ([]models.Location, error) {
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

func fetchLocationByID(ID int) (*models.Relation, error) {
	relations, err := FetchRelations()
	if err != nil {
		return nil, err
	}
	for i := range relations {
		if relations[i].ID == ID {
			return &relations[i], nil
		}
	}
	return nil, errors.New("location not found")
}
