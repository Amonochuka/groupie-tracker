# Groupie Tracker

Groupie Tracker is a web application written in Go that displays information about music artists and bands using data from the Groupie Tracker API.

The application allows users to browse artists, view their members, and explore concert locations and dates.

The goal of this project is to practice backend development in Go, including:

 - Working with external APIs

 - JSON decoding

 - Struct modelling

 - HTTP servers and handlers

 - Template rendering

 - Testing with mocks

## Features

 - View all artists from the API

 - Search artists by band name

 - Search artists by member name

 - View artist details

 - Display concert locations and dates

Clean separation between:

 - API layer

 - Handlers

 - Models

 - Templates

Unit tests using mock HTTP clients

### Technologies Used

 - Go (Golang)

 - HTML templates

 - Groupie Tracker API

 - net/http

 - encoding/json

 - Go testing framework

#### API Used

Data is retrieved from the public API:

https://groupietrackers.herokuapp.com/api

Endpoints used:

 - /artists

 - /locations

 - /dates

 - /relation

```
Project Structure
groupie-tracker
│
├── api
│   ├── fetchArtists.go
│   ├── fetchLocations.go
│   ├── fetchDates.go
│   ├── fetchRelations.go
│   └── *_test.go
│
├── handlers
│   ├── home.go
│   └── artist.go
│
├── models
│   └── models.go
│
├── templates
│   ├── index.html
│   └── artist.html
│
├── static
│
├── main.go
└── README.md
```

## Installation

 - Clone the repository:

git clone https://github.com/yourusername/groupie-tracker.git

 - Navigate to the project:

cd groupie-tracker

 - Run the application:

go run main.go

The server will start at:

http://localhost:8080

### Search Functionality

The application supports searching by:

 - Artist/Band name

 - Member name

Example:

?search=metallica

or

?search=james

### Testing
Use ```go test -v -count=1 ./... ``` to check all test files without caching

Tests are implemented for the API layer using mock HTTP clients.

This ensures:

 - HTTP requests are correctly formed

 - JSON responses are properly decoded

 - Struct mappings are correct

Run tests with:

```
go test ./...
```

Example tested functions:

 - FetchArtists

 - FetchLocations

 - FetchDates

 - FetchRelations

***Mock clients simulate API responses to allow unit testing without calling the real API.***

### Concepts Practiced

This project demonstrates several important Go concepts:

 - HTTP handlers

 - Template parsing and execution

 - Struct composition

 - JSON decoding

 - Dependency injection

 - Mock testing

 - Slice and pointer behavior

 - Error handling

### Example Workflow

1. The server receives a request

GET /

2.  The handler calls the API layer

FetchArtists()
FetchRelations()

3. Data is combined into a view struct

ArtistView

4. The template engine renders the HTML page

index.html

5. The final HTML is sent to the browser
 
***Possible Improvements***

Future improvements could include:

 - Caching API responses

 - Better frontend styling

 - Adding a REST API layer

- Implementing concurrency for API requests

### Learning Goals

This project was built to improve understanding of:

 - Backend architecture in Go

 - Clean project structure

 - API integration

 - Testing strategies

### Authors

Flovian Atieno

Stephen Jarso

Amon Ochuka
