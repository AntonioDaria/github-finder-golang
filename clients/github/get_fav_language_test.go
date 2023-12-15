package github

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFavLanguage(t *testing.T) {
	// create test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// write response
		response := []Repo{
			{
				Language: "Go",
			},
		}

		// check request method
		if r.Method != http.MethodGet {
			t.Errorf("Expected 'GET' request, got '%s'", r.Method)
		}

		// check request url
		if r.URL.String() != "/users/githubUserName/repos" {
			t.Errorf("Expected request to '/users/githubUserName/repos', got '%s'", r.URL.String())
		}

		// send response to be tested
		if r.URL.Path == "/users/githubUserName/repos" {
			w.WriteHeader(http.StatusOK)
		}

		// check response
		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()

	// create new client
	client := NewClient(http.DefaultClient, ts.URL)

	// make request
	language, err := client.GetFavLanguage(context.Background(), "githubUserName")
	assert.NoError(t, err)

	// assert response
	assert.Equal(t, "Go", language)
}
