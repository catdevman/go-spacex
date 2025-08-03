package spacex

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

// setup sets up a test HTTP server along with a spacex.Client that is
// configured to talk to that test server. Tests should register handlers on
// mux which provide mock responses for the API method being tested.
func setup() (client *Client, mux *http.ServeMux, serverURL string, teardown func()) {
	// mux is the HTTP request multiplexer used with the test server.
	mux = http.NewServeMux()

	// server is a test HTTP server used to provide mock API responses.
	server := httptest.NewServer(mux)

	// client is the SpaceX client being tested and is configured to use that server.
	client = NewClient(nil)
	url, _ := url.Parse(server.URL + "/")
	client.BaseURL = url

	return client, mux, server.URL, server.Close
}

func TestDragonsService_GetDragon(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/dragons/5e9d058759b1ff74a7ad5f8f", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Request method = %v, want %v", r.Method, "GET")
		}
		fmt.Fprint(w, `{"id":"5e9d058759b1ff74a7ad5f8f","name":"Dragon 1"}`)
	})

	ctx := context.Background()
	dragon, err := client.Dragons.GetDragon(ctx, "5e9d058759b1ff74a7ad5f8f")
	if err != nil {
		t.Fatalf("Dragons.GetDragon returned error: %v", err)
	}

	want := &Dragon{ID: "5e9d058759b1ff74a7ad5f8f", Name: "Dragon 1"}
	if dragon.Name != want.Name {
		t.Errorf("Dragons.GetDragon returned %+v, want %+v", dragon.Name, want.Name)
	}
}
