package poker

import (
	"io"
	"time"
)

type TexasHoldem struct {
	alerter BlindAlerter
	store   PlayerStore
}

func NewGame(store PlayerStore, alerter BlindAlerter) *TexasHoldem {
	return &TexasHoldem{store: store, alerter: alerter}
}

func (g *TexasHoldem) Start(numberOfPlayers int, alertsDestination io.Writer) {
	blindIncrement := time.Duration(5+numberOfPlayers) * time.Minute
	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second

	for _, blind := range blinds {
		g.alerter.ScheduleAlertAt(blindTime, blind, alertsDestination)
		blindTime = blindTime + blindIncrement
	}
}
func (g *TexasHoldem) Over(winner string) {
	g.store.RecordWin(winner)
}
