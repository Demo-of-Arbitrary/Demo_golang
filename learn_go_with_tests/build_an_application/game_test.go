package poker_test

import (
	"fmt"
	"testing"
	"time"

	poker "github.com/Demo_golang/learn_go_with_tests/build_an_application"
)

func TestGame_Start(t *testing.T) {
	t.Run("schedules alerts on game start for 5 players", func(t *testing.T) {
		blindAlerter := &SpyBlindAlerter{}
		game := poker.NewGame(dummyPlayerStore, blindAlerter)

		game.Start(5)
		cases := []ScheduledAlerter{
			{At: 0 * time.Second, Amount: 100},
			{At: 10 * time.Minute, Amount: 200},
			{At: 20 * time.Minute, Amount: 300},
			{At: 30 * time.Minute, Amount: 400},
			{At: 40 * time.Minute, Amount: 500},
			{At: 50 * time.Minute, Amount: 600},
			{At: 60 * time.Minute, Amount: 800},
			{At: 70 * time.Minute, Amount: 1000},
			{At: 80 * time.Minute, Amount: 2000},
			{At: 90 * time.Minute, Amount: 4000},
			{At: 100 * time.Minute, Amount: 8000},
		}

		checkSchedulingCases(cases, t, blindAlerter)
	})

	t.Run("schedules alerts on game start for 7 players", func(t *testing.T) {
		blindAlerter := &SpyBlindAlerter{}
		game := poker.NewGame(dummyPlayerStore, blindAlerter)

		game.Start(7)
		cases := []ScheduledAlerter{
			{At: 0 * time.Second, Amount: 100},
			{At: 12 * time.Minute, Amount: 200},
			{At: 24 * time.Minute, Amount: 300},
			{At: 36 * time.Minute, Amount: 400},
		}
		checkSchedulingCases(cases, t, blindAlerter)
	})
}

func TestGame_Finish(t *testing.T) {
	store := &poker.StubPlayerStore{}
	game := poker.NewGame(store, dummyBlindAlerter)
	winner := "Ruth"

	game.Over(winner)
	poker.AssertPlayerWin(t, store, winner)
}

func checkSchedulingCases(cases []ScheduledAlerter, t *testing.T, blindAlerter *SpyBlindAlerter) {
	for i, want := range cases {
		t.Run(fmt.Sprint(want), func(t *testing.T) {
			if len(blindAlerter.alerts) <= i {
				t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.alerts)
			}

			got := blindAlerter.alerts[i]
			assertScheduledAlert(t, got, want)
		})
	}
}

func assertScheduledAlert(t *testing.T, got, c ScheduledAlerter) {
	amountGot := got.Amount
	if amountGot != c.Amount {
		t.Errorf("got amount %d, want %d", amountGot, c.Amount)
	}

	gotScheduledTime := got.At
	if gotScheduledTime != c.At {
		t.Errorf("got scheduled time of %v, want %v", gotScheduledTime, c.At)
	}
}
