package api
import(
	"net/http"
	"encoding/json"
	"groupie-tracker/models"
)
var artists []models.Artist
func FetchArtists()error{
resp,err:=http.Get("https://groupietrackers.herokuapp.com/api/artists" )
	if err !=nil{
		return err
	}
defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(&artists)
}
