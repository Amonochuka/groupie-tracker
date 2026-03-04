
package handlers

import(
	"net/http"
	"html/template"
	"groupie-tracker/models"
)
 var artists []models.Artist

func HomeHandler(w http.ResponseWriter, r *http.Request){
	templ,err:=template.ParseFiles("templates/index.html")
	if err != nil{
		http.Error(w,"Template error",500)
	}
	templ.Execute(w,artists)
}