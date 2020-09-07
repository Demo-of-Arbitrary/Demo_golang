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

type CountdownOperationSpy struct {
	Calls []string
}

func (c *CountdownOperationSpy) Sleep() {
	c.Calls = append(c.Calls, sleep)
}

func (c *CountdownOperationSpy) Write(p []byte) (n int, err error) {
	c.Calls = append(c.Calls, writer)
	return
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(d time.Duration) {
	s.durationSlept = d
}

const writer = "writer"
const sleep = "sleep"

const countDownNum = 3
const finalValue = "Go!"

func Countdown(b io.Writer, sleeper Sleeper) {
	for i := countDownNum; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(b, i)
	}

	sleeper.Sleep()
	fmt.Fprint(b, finalValue)
}

func main() {
	sleeper := ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, &sleeper)
}
