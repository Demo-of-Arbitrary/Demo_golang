package poker_test

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"
	"time"

	poker "github.com/Demo_golang/learn_go_with_tests/build_an_application"
)

var dummyBlindAlerter = new(SpyBlindAlerter)
var dummyPlayerStore = new(poker.StubPlayerStore)
var dummyStdout = new(bytes.Buffer)

type SpyBlindAlerter struct {
	alerts []ScheduledAlerter
}
type ScheduledAlerter struct {
	At     time.Duration
	Amount int
	To     io.Writer
}

func (s *ScheduledAlerter) String() string {
	return fmt.Sprintf("%d chips at %v", s.Amount, s.At)
}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int, to io.Writer) {
	s.alerts = append(s.alerts, ScheduledAlerter{
		duration, amount, to,
	})
}

func TestCLI(t *testing.T) {
	t.Run("record chris win form user input", func(t *testing.T) {
		in := strings.NewReader("5\nChris wins\n")

		game := &poker.GameSpy{}

		cli := poker.NewCLI(in, dummyStdout, game)
		cli.PlayPoker()

		if game.FinshedWith != "Chris" {
			t.Errorf("expected finished with %s, got %s", "Chris", game.FinshedWith)
		}
	})

	t.Run("record cleo win form user input", func(t *testing.T) {
		in := strings.NewReader("5\nCleo wins\n")

		game := &poker.GameSpy{}

		cli := poker.NewCLI(in, dummyStdout, game)
		cli.PlayPoker()

		if game.FinshedWith != "Cleo" {
			t.Errorf("expected finished with %s, got %s", "Cleo", game.FinshedWith)
		}
	})

	t.Run("it prompts the user to enter the number of players and starts the game", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("7\n")
		game := &poker.GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertMessageSentToUser(t, stdout, poker.PlayerPrompt)

		if game.StartedWith != 7 {
			t.Errorf("wanted Start called with 7 but got %d", game.StartedWith)
		}
	})

	t.Run("it prints an error when a non numeric value is entered and doesn't start the game", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("Pies\n")
		game := &poker.GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertMessageSentToUser(t, stdout, poker.PlayerPrompt, poker.BadPlayerInputErrMsg)

		if game.StartCalled {
			t.Errorf("game should not have started")
		}
	})

	t.Run("it prints an error when not specified end statement given and game should not have ended", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("7\nLloyd is a killer\n")
		game := &poker.GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		if !game.StartCalled {
			t.Errorf("game should have started")
		}
		if game.OverCalled {
			t.Errorf("game should have not ended")
		}
	})
}

func assertMessageSentToUser(t *testing.T, stdout *bytes.Buffer, messages ...string) {
	t.Helper()
	want := strings.Join(messages, "")
	got := stdout.String()
	if got != want {
		t.Errorf("got %q sent to stdout but expected %q", got, messages)
	}
}
