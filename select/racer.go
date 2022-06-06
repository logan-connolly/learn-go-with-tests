package racer

import (
	"errors"
	"net/http"
	"time"
)

var ErrMissingArgs = errors.New("Not enough args provided.")

type result struct {
	url  string
	time time.Duration
}

func Racer(urls ...string) (string, error) {
	if len(urls) < 2 {
		return "", ErrMissingArgs
	}
	rc := make(chan result)

	for _, url := range urls {
		go func(u string) {
			start := time.Now()
			http.Get(u)
			duration := time.Since(start)
			rc <- result{url: u, time: duration}
		}(url)
	}

	r := <-rc
	best := result{r.url, r.time}

	for i := 1; i < len(urls); i++ {
		r := <-rc
		if r.time < best.time {
			best = r
		}
	}

	return best.url, nil
}
