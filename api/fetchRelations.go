package api

import (
	"encoding/json"
	"errors"
	"groupie-tracker/models"
)

func FetchRelations() ([]models.Relation, error) {
	resp, err := DefaultClient.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data models.RelationResponse

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	return data.Index, nil

}

func fetchRelationByID(ID int) (*models.Location, error) {
	locations, err := FetchLocations()
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
