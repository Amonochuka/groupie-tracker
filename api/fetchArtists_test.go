// api/fetchArtists_test.go
package api

import (
	"bytes"
	"io"
	"net/http"
	"testing"
)

type MockClient struct {
	Response *http.Response
	Err      error
}

// implement Do() for MockClient
func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return m.Response, m.Err
}

func TestFetchArtists(t *testing.T) {
	mockJSON := `[
		{
			"id": 1,
			"name": "Band 1",
			"image": "https://example.com/band1.jpg",
			"members": ["Alice", "Bob"],
			"creationDate": 2001,
			"firstAlbum": "First Sound",
			"concertDates": "https://api.groupietrackers.com/api/dates/1",
			"locations": "https://api.groupietrackers.com/api/locations/1",
			"relations": "https://api.groupietrackers.com/api/relation/1"
		},
		{
			"id": 2,
			"name": "Band 2",
			"image": "https://example.com/band2.jpg",
			"members": ["Charlie", "Dave", "Eve"],
			"creationDate": 1998,
			"firstAlbum": "Echo Waves",
			"concertDates": "https://api.groupietrackers.com/api/dates/2",
			"locations": "https://api.groupietrackers.com/api/locations/2",
			"relations": "https://api.groupietrackers.com/api/relation/2"
		}
	]`
	mockResp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewBufferString(mockJSON)),
	}

	client := &MockClient{Response: mockResp}
	artists, err := FetchArtists(client)
	if err != nil {
		t.Fatal(err)
	}

	if len(artists) != 1 || artists[0].Name != "Band 1" {
		t.Errorf("expected 1 artist with name Band 1, got %+v", artists)
	}
}