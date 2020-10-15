package poker_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	poker "github.com/Demo_golang/learn_go_with_tests/build_an_application"
)

type SpyBlindAlerter struct {
	alerts []scheduledAlerter
}
type scheduledAlerter struct {
	at     time.Duration
	amount int
}

func (s *scheduledAlerter) String() string {
	return fmt.Sprintf("%d chips at %v", s.amount, s.at)
}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	s.alerts = append(s.alerts, scheduledAlerter{
		duration, amount,
	})
}

func TestCLI(t *testing.T) {
	t.Run("record chris win form user input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := &poker.StubPlayerStore{}
		dummySpyAlterer := &SpyBlindAlerter{}

		cli := poker.NewCLI(playerStore, in, dummySpyAlterer)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Chris")
	})

	t.Run("record cleo win form user input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		playerStore := &poker.StubPlayerStore{}
		dummySpyAlterer := &SpyBlindAlerter{}

		cli := poker.NewCLI(playerStore, in, dummySpyAlterer)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Cleo")
	})

	t.Run("it schedules printing of blind values", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := &poker.StubPlayerStore{}
		blindAlerter := &SpyBlindAlerter{}

		cli := poker.NewCLI(playerStore, in, blindAlerter)
		cli.PlayPoker()

		cases := []scheduledAlerter{
			{0 * time.Second, 100},
			{10 * time.Minute, 200},
			{20 * time.Minute, 300},
			{30 * time.Minute, 400},
			{40 * time.Minute, 500},
			{50 * time.Minute, 600},
			{60 * time.Minute, 800},
			{70 * time.Minute, 1000},
			{80 * time.Minute, 2000},
			{90 * time.Minute, 4000},
			{100 * time.Minute, 8000},
		}

		for i, c := range cases {
			t.Run(fmt.Sprintf("%d scheduled for %v", c.amount, c.at), func(t *testing.T) {

				if len(blindAlerter.alerts) <= i {
					t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.alerts)
				}

				got := blindAlerter.alerts[i]
				assertScheduledAlert(t, got, c)

			})
		}

	})
}

func assertScheduledAlert(t *testing.T, got, c scheduledAlerter) {

	amountGot := got.amount
	if amountGot != c.amount {
		t.Errorf("got amount %d, want %d", amountGot, c.amount)
	}

	gotScheduledTime := got.at
	if gotScheduledTime != c.at {
		t.Errorf("got scheduled time of %v, want %v", gotScheduledTime, c.at)
	}
}
