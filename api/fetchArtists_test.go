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
	mockJSON := `[{"ID":1,"Name":"Band 1"}]`
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
