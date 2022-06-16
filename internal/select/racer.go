package racer

import (
	"errors"
	"net/http"
	"time"
)

var ErrRacerTimeout = errors.New("Race took too long.")

type RacePair struct {
	urlOne string
	urlTwo string
}

func Racer(rp RacePair, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(rp.urlOne):
		return rp.urlOne, nil
	case <-ping(rp.urlTwo):
		return rp.urlTwo, nil
	case <-time.After(timeout):
		return "", ErrRacerTimeout
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
