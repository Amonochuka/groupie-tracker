package api

import (
	"errors"
	"groupie-tracker/models"

	"encoding/json"
	"net/http"
)

func GetDates() ([]models.Date, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/dates")
	if err != nil {
		return nil, err
	}
	var data models.DateResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	return data.Index, err
}

func getDateByID(ID int) (*models.Date, error) {
	dates, err := GetDates()
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
