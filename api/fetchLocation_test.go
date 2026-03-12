// api/fetchLocations_test.go
package api

import (
	"bytes"
	"io"
	"net/http"
	"testing"
)

// TestGetLocations tests FetchLocations with a mocked HTTP response
func TestGetLocations(t *testing.T) {
	// mock JSON response
	mockJSON := `{
		"index": [
			{
				"id": 1,
				"locations": ["new_york-usa", "london-uk"],
				"dates": "2024-03-01"
			},
			{
				"id": 2,
				"locations": ["paris-france"],
				"dates": "2024-04-10"
			}
		]
	}`

	// fake HTTP response
	mockResp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewBufferString(mockJSON)),
	}

	// mocked client
	client := &MockClient{Response: mockResp}

	// call FetchLocations
	locations, err := FetchLocations(client)
	if err != nil {
		t.Fatal(err)
	}

	// assertions
	if len(locations) != 2 {
		t.Errorf("expected 2 locations, got %d", len(locations))
	}

	if locations[0].ID != 1 || locations[1].ID != 2 {
		t.Errorf("unexpected IDs: %+v", locations)
	}

	if len(locations[0].Locations) != 2 {
		t.Errorf("unexpected locations for ID 1: %+v", locations[0].Locations)
	}

	if locations[0].Locations[0] != "new_york-usa" {
		t.Errorf("unexpected first location: %s", locations[0].Locations[0])
	}

	if locations[0].Dates != "2024-03-01" {
		t.Errorf("unexpected date: %s", locations[0].Dates)
	}
}
