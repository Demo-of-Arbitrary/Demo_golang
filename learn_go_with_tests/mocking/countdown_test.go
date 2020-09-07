package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestCountdown(t *testing.T) {
	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}

		Countdown(buffer, &CountdownOperationSpy{})

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("sleep before every print", func(t *testing.T) {
		spySleeperPrint := &CountdownOperationSpy{}
		Countdown(spySleeperPrint, spySleeperPrint)

		want := []string{
			sleep,
			writer,
			sleep,
			writer,
			sleep,
			writer,
			sleep,
			writer,
		}

		if !reflect.DeepEqual(want, spySleeperPrint.Calls) {
			t.Errorf("wanted calls %+v got %+v", want, spySleeperPrint.Calls)
		}
	})
}
