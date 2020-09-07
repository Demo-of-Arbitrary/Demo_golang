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

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

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
