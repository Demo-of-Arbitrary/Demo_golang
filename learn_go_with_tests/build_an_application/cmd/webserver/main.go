package main

import (
	"log"
	"net/http"
	"os"

	poker "github.com/Demo_golang/learn_go_with_tests/build_an_application"
)

const dbFileName = "game.db.json"

func main() {
	database, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem opening database: %s %v", dbFileName, err)
	}

	store, err := poker.NewFileSystemPlayerStore(database)

	if err != nil {
		log.Fatalf("problem creating file system player store: %v", err)
	}

	game := poker.NewGame(store, poker.BlindAlerterFunc(poker.Alerter))
	server, err := poker.NewPlayerServer(store, game)
	if err != nil {
		log.Fatalf("problem creating server: %v", err)
	}

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
