package poker_test

import (
	"bytes"
	"fmt"
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
}

func (s *ScheduledAlerter) String() string {
	return fmt.Sprintf("%d chips at %v", s.Amount, s.At)
}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	s.alerts = append(s.alerts, ScheduledAlerter{
		duration, amount,
	})
}

type GameSpy struct {
	StartedWith int
	FinshedWith string
}

func (g *GameSpy) Start(numberOfPlayers int) {
	g.StartedWith = numberOfPlayers
}
func (g *GameSpy) Over(winner string) {
	g.FinshedWith = winner
}

func TestCLI(t *testing.T) {
	t.Run("record chris win form user input", func(t *testing.T) {
		in := strings.NewReader("5\nChris wins\n")

		game := &GameSpy{}

		cli := poker.NewCLI(in, dummyStdout, game)
		cli.PlayPoker()

		if game.FinshedWith != "Chris" {
			t.Errorf("expected finished with %s, got %s", "Chris", game.FinshedWith)
		}
	})

	t.Run("record cleo win form user input", func(t *testing.T) {
		in := strings.NewReader("5\nCleo wins\n")

		game := &GameSpy{}

		cli := poker.NewCLI(in, dummyStdout, game)
		cli.PlayPoker()

		if game.FinshedWith != "Cleo" {
			t.Errorf("expected finished with %s, got %s", "Cleo", game.FinshedWith)
		}
	})

	t.Run("it prompts the user to enter the number of players and starts the game", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("7\n")
		game := &GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		gotPrompt := stdout.String()
		wantPrompt := poker.PlayerPrompt

		if gotPrompt != wantPrompt {
			t.Errorf("got %q, want %q", gotPrompt, wantPrompt)
		}

		if game.StartedWith != 7 {
			t.Errorf("wanted Start called with 7 but got %d", game.StartedWith)
		}
	})
}
