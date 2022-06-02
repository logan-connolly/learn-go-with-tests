package main

import (
	"bytes"
	"testing"
	"time"
)

var spyCounter = 0

func mockSleep(d time.Duration) {
	spyCounter += 1
}

func TestCountdownClock(t *testing.T) {
	buffer := &bytes.Buffer{}
	sleeper := &ConfigurableSleeper{
		iterations: 3,
		duration:   time.Second * 1,
		sleep:      mockSleep,
	}
	clock := CountdownClock{
		writer:  buffer,
		sleeper: *sleeper,
	}
	clock.Start()

	got := buffer.String()
	want := "3\n2\n1\nGo!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	if spyCounter != 3 {
		t.Errorf("expected 3 calls got %v", spyCounter)
	}
}
