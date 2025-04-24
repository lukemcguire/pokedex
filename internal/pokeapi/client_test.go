package pokeapi

import (
	"bytes"
	"io"
	"net/http" // Added for deep comparison of results
	"testing"  // Added missing import
)

// MockTransport is a custom http.RoundTripper for testing
type MockTransport struct {
	// Map request URLs to mock responses
	responses map[string]MockResponse
}

type MockResponse struct {
	StatusCode int
	Body       string
	Err        error
}

// RoundTrip implements the http.RoundTripper interface
func (m *MockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	mockResp, exists := m.responses[req.URL.String()]
	if !exists {
		return &http.Response{
			StatusCode: 404,
			Body:       io.NopCloser(bytes.NewBufferString("not found")),
			Header:     make(http.Header), // Added Header for completeness
		}, nil
	}

	if mockResp.Err != nil {
		return nil, mockResp.Err
	}

	return &http.Response{
		StatusCode: mockResp.StatusCode,
		Body:       io.NopCloser(bytes.NewBufferString(mockResp.Body)),
		Header:     make(http.Header),
	}, nil
}
func TestListLocationAreas(t *testing.T) {
	// Test case 1: Successful fetch of first page
	t.Run("FetchFirstPage", func(t *testing.T) {
		// Create sample JSON response
		mockLocationsJSON := `{
            "count": 2,
            "next": "https://pokeapi.co/api/v2/location-area?offset=20&limit=20",
            "previous": null,
            "results": [
                {
                    "name": "test-location-1",
                    "url": "https://pokeapi.co/api/v2/location-area/1/"
                },
                {
                    "name": "test-location-2",
                    "url": "https://pokeapi.co/api/v2/location-area/2/"
                }
            ]
        }`

		// Setup mock transport
		mockTransport := &MockTransport{
			responses: map[string]MockResponse{
				// Use BaseURL constant if defined, otherwise hardcode
				// Assuming BaseURL = "https://pokeapi.co/api/v2"
				"https://pokeapi.co/api/v2/location-area": { // Use BaseURL + path
					StatusCode: 200,
					Body:       mockLocationsJSON,
				},
			},
		}

		// Create client with mock transport
		httpClient := &http.Client{Transport: mockTransport}
		// Assuming NewClient exists and accepts functional options
		// Also assuming a Client struct exists with httpClient field
		client := NewClient(func(c *Client) {
			c.httpClient = httpClient
		})

		// Make the API call
		// Assuming ListLocationAreas takes an optional *string for the URL
		// and returns a struct similar to LocationAreasResponse and an error
		result, err := client.ListLocationAreas(nil) // Pass nil for the first page

		// Assertions
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if result.Count != 2 {
			t.Errorf("Expected count 2, got %d", result.Count)
		}

		expectedNextURL := "https://pokeapi.co/api/v2/location-area?offset=20&limit=20"
		if result.Next == nil || *result.Next != expectedNextURL {
			t.Errorf("Expected next URL %q, got %v", expectedNextURL, result.Next)
		}

		if result.Previous != nil {
			t.Errorf("Expected previous URL to be nil, got %v", result.Previous)
		}

		// Define expected results (assuming the result struct has a similar structure)
		// You'll need to adjust 'struct{ Name string; URL string }' if your actual
		// result struct field type is different (e.g., a named struct like 'LocationAreaResult')
		expectedResults := []struct {
			Name string
			URL  string
		}{
			{Name: "test-location-1", URL: "https://pokeapi.co/api/v2/location-area/1/"},
			{Name: "test-location-2", URL: "https://pokeapi.co/api/v2/location-area/2/"},
		}

		// Check length first
		if len(result.Results) != len(expectedResults) {
			t.Fatalf("Expected %d results, got %d", len(expectedResults), len(result.Results))
		}

	}) // <-- Added missing closing brace for t.Run

	// --- Add more test cases here ---
	// t.Run("FetchNextPage", func(t *testing.T) { ... })
	// t.Run("HandleAPIError", func(t *testing.T) { ... })
	// t.Run("HandleNotFound", func(t *testing.T) { ... })

} // <-- Added missing closing brace for TestListLocationAreas
