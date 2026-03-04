package api
import(
	"net/http"
	"encoding/json"
	"groupie-tracker/models"
)
var artists []models.Artist

//retrieve data from the Api and decode it to artists slice
func FetchArtists() ([]models.Artist, error){
resp,err:=http.Get("https://groupietrackers.herokuapp.com/api/artists" )
	if err !=nil{
		return nil, err
	}
defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&artists)
	return artists ,err
}
