// api/fetchDates_test.go
package api

import (
	"bytes"
	"io"
	"net/http"
	"testing"
)

// TestGetDates tests FetchDates with a mocked HTTP response
func TestGetDates(t *testing.T) {
	// mock JSON for DateResponse
	mockJSON := `{
		"index": [
			{
				"id": 1,
				"dates": ["2024-03-01", "2024-03-02"]
			},
			{
				"id": 2,
				"dates": ["2024-04-10"]
			}
		]
	}`

	// fake HTTP response
	mockResp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewBufferString(mockJSON)),
	}

	// mock client
	client := &MockClient{Response: mockResp}

	// call FetchDates
	dates, err := GetDates(client)
	if err != nil {
		t.Fatal(err)
	}

	// assertions
	if len(dates) != 2 {
		t.Errorf("expected 2 dates, got %d", len(dates))
	}

	if dates[0].ID != 1 || dates[1].ID != 2 {
		t.Errorf("unexpected IDs: %+v", dates)
	}

	if len(dates[0].Dates) != 2 {
		t.Errorf("unexpected number of dates for ID 1: %+v", dates[0].Dates)
	}

	if dates[0].Dates[0] != "2024-03-01" {
		t.Errorf("unexpected first date: %s", dates[0].Dates[0])
	}
}
