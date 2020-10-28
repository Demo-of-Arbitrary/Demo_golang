package main

import (
	"fmt"
	"log"
	"os"

	poker "github.com/Demo_golang/learn_go_with_tests/build_an_application"
)

const dbFileName = "game.db.json"

func main() {
	fmt.Println("let's play poker")
	fmt.Println("Type {Name} wins to record a win")

	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}

	store, err := poker.NewFileSystemPlayerStore(db)

	if err != nil {
		log.Fatalf("problem creating file system player store %v", err)
	}

	game := poker.NewGame(store, poker.BlindAlerterFunc(poker.Alerter))
	cli := poker.NewCLI(os.Stdin, os.Stdout, game)
	cli.PlayPoker()
}
