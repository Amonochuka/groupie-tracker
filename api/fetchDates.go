package api

import (
	"groupie-tracker/models"

	"net/http"
	"encoding/json"
)
func GetDates()([]models.Date,error){
	resp,err:=http.Get("https://groupietrackers.herokuapp.com/api/dates")
if err != nil{
	return nil ,err
}
var data models.DateResponse
err=json.NewDecoder(resp.Body).Decode(&data)
return data.Index, err
}