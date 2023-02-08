package countdown

import (
	"fmt"
	"io"
	"os"
	"time"
)

type ConfigurableSleeper struct {
	iterations int
	duration   time.Duration
	sleep      func(time.Duration)
}

func (s *ConfigurableSleeper) Sleep() {
	s.sleep(s.duration)
}

type CountdownClock struct {
	writer  io.Writer
	sleeper ConfigurableSleeper
}

func (c *CountdownClock) Start() {
	for i := c.sleeper.iterations; i > 0; i-- {
		fmt.Fprintln(c.writer, i)
		c.sleeper.Sleep()
	}

	fmt.Fprint(c.writer, "Go!")
}

func main() {
	sleeper := &ConfigurableSleeper{
		iterations: 3,
		duration:   time.Second,
		sleep:      time.Sleep,
	}
	c := CountdownClock{
		writer:  os.Stdout,
		sleeper: *sleeper,
	}
	c.Start()
}
