package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (s *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

type CountdownClock struct {
	writer         io.Writer
	sleeper        Sleeper
	startingNumber int
}

func (c *CountdownClock) Start() {
	for i := c.startingNumber; i > 0; i-- {
		c.sleeper.Sleep()
		fmt.Fprintln(c.writer, i)
	}

	c.sleeper.Sleep()
	fmt.Fprint(c.writer, "Go!")
}

func main() {
	c := CountdownClock{
		writer:         os.Stdout,
		sleeper:        &DefaultSleeper{},
		startingNumber: 3,
	}
	c.Start()
}
