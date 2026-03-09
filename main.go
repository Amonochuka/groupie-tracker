package main

import (
	"fmt"
	"log"
	"net/http"
	// "strconv"

	"groupie-tracker/api"
	"groupie-tracker/handlers"


)


func main(){
	artists, err:=api.FetchArtists()
	fmt.Println(len(artists))
	if err!=nil{
		log.Fatal(err)
	}
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/",handlers.HomeHandler)
	http.HandleFunc("/artist",handlers.ArtistHandler)
	log.Println("Server running on http://localhost:8080")
	log.Println(http.ListenAndServe(":8080",nil))

}
