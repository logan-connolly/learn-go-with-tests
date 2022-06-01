package main

import (
	"bytes"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCountdownClock(t *testing.T) {
	buffer := &bytes.Buffer{}
	sleeper := &SpySleeper{}
	clock := CountdownClock{
		writer:         buffer,
		sleeper:        sleeper,
		startingNumber: 3,
	}

	clock.Start()

	got := buffer.String()
	want := "3\n2\n1\nGo!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	if sleeper.Calls != 4 {
		t.Errorf("expected 4 calls got %q", sleeper.Calls)
	}
}
