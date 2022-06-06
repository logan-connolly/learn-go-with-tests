package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	slowServer := createMockServer(20 * time.Millisecond)
	fastServer := createMockServer(1 * time.Millisecond)

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	want := fastURL
	got, err := Racer(slowURL, fastURL)

	defer slowServer.Close()
	defer fastServer.Close()

	if err != nil {
		t.Fatal("Did not expect to get an err", err)
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func createMockServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(delay)
			w.WriteHeader(http.StatusOK)
		}))
}
