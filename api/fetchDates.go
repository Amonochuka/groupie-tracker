package api

import (
	"errors"
	"groupie-tracker/models"
	"net/http"

	"encoding/json"
)

func GetDates(client HTTPClient) ([]models.Date, error) {
	req, _ := http.NewRequest("GET", "https://groupietrackers.herokuapp.com/api/dates", nil)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	var data models.DateResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	return data.Index, err
}

func getDateByID(client HTTPClient, ID int) (*models.Date, error) {
	dates, err := GetDates(client)
	if err != nil {
		return nil, err
	}

	for i := range dates {
		if dates[i].ID == ID {
			return &dates[i], nil
		}
	}
	return nil, errors.New(("date not found"))
}
