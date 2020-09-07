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

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
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
	sleeper := DefaultSleeper{}
	Countdown(os.Stdout, &sleeper)
}
