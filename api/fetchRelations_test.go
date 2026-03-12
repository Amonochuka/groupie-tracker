// api/fetchRelations_test.go
package api

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"testing"
)

// TestFetchRelations tests FetchRelations with a mocked HTTP response
func TestFetchRelations(t *testing.T) {
	// mock JSON for RelationResponse
	mockJSON := `{
		"index": [
			{
				"ID": 1,
				"DatesLocations": {
					"new_york-usa": ["2024-03-01"],
					"london-uk": ["2024-03-05"]
				}
			},
			{
				"ID": 2,
				"DatesLocations": {
					"paris-france": ["2024-04-10"]
				}
			}
		]
	}`

	// wrap it in a fake http.Response
	mockResp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewBufferString(mockJSON)),
	}

	// use MockClient instead of real http.Client
	client := &MockClient{Response: mockResp}

	// call your FetchRelations function
	relations, err := FetchRelations(client)
	if err != nil {
		t.Fatal(err)
	}

	// assert results
	if len(relations) != 2 {
		t.Errorf("expected 2 relations, got %d", len(relations))
	}

	if relations[0].ID != 1 || relations[1].ID != 2 {
		t.Errorf("unexpected IDs: %+v", relations)
	}

	if len(relations[0].DatesLocations["new_york-usa"]) != 1 {
		t.Errorf("unexpected dates for new_york-usa: %+v", relations[0].DatesLocations)
	}
	fmt.Println(len(relations))
}
