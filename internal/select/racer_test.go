package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("returns faster server", func(t *testing.T) {
		slowServer := createMockServer(20 * time.Millisecond)
		fastServer := createMockServer(1 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		rp := RacePair{slowServer.URL, fastServer.URL}
		timeout := 30 * time.Millisecond

		want := fastServer.URL
		got, err := Racer(rp, timeout)
		if err != nil {
			t.Fatalf("did not expect to receive error, but got %v", err)
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns error if timeout reached", func(t *testing.T) {
		slowServer := createMockServer(11 * time.Millisecond)
		slowerServer := createMockServer(12 * time.Millisecond)

		defer slowServer.Close()
		defer slowerServer.Close()

		rp := RacePair{slowServer.URL, slowerServer.URL}
		timeout := 10 * time.Millisecond

		_, err := Racer(rp, timeout)

		if err != ErrRacerTimeout {
			t.Errorf("expected to get ErrRacerTimeout, but got %v", err)
		}
	})
}

func createMockServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(delay)
			w.WriteHeader(http.StatusOK)
		}))
}
