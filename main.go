package main

import (
	// "fmt"
	"log"
	"net/http"
	// "strconv"

	"groupie-tracker/api"
	"groupie-tracker/handlers"


)


func main(){
	err:=api.FetchArtists()
	// fmt.Println(len(artists))
	if err!=nil{
		log.Fatal(err)
	}
	http.HandleFunc("/",handlers.HomeHandler)
	log.Println("Server running on http://localhost:8080")
	log.Println(http.ListenAndServe(":8080",nil))

}
