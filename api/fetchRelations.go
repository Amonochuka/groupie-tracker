package api

import (
	"encoding/json"
	"errors"
	"groupie-tracker/models"
	"net/http"
)

func FetchRelations(client HTTPClient) ([]models.Relation, error) {
	req, _ := http.NewRequest("GET", "https://groupietrackers.herokuapp.com/api/relation", nil)
	resp, err := client.Do(req)
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

func fetchRelationByID(client HTTPClient, ID int) (*models.Relation, error) {
	relations, err := FetchRelations(client)
	if err != nil {
		return nil, err
	}
	for i := range relations {
		if relations[i].ID == ID {
			return &relations[i], nil
		}
	}
	return nil, errors.New("relation not found")
}
