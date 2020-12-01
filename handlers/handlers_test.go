package handlers

import (
	"testing"
	"net/http/httptest"
	"net/http"
)

func TestHealthCheckHandler(t *testing.T) {
	server := httptest.NewServer(Handlers())
	defer server.Close()

	resp, err := http.Get(server.URL + "/health")

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Servers didn't return HTTP 200.Response Code: %d", resp.StatusCode);
	}
}
