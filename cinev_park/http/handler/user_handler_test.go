package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMovie(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/movies",Movies)
	testServ := httptest.NewTLSServer(mux)
	defer testServ.Close()

	testClient := testServ.Client()
	url := testServ.URL

	resp, err := testClient.Get(url + "/movies")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d; got %d", http.StatusOK, resp.StatusCode)
	}
	defer resp.Body.Close()

}
