package poker

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Game interface {
	Start(n int)
	Over(w string)
}

type CLI struct {
	in   *bufio.Scanner
	out  io.Writer
	game Game
}

func NewCLI(in io.Reader, out io.Writer, game Game) *CLI {
	return &CLI{in: bufio.NewScanner(in), out: out, game: game}
}

const PlayerPrompt = "Please enter the number of players: "

func (c *CLI) PlayPoker() {
	fmt.Fprint(c.out, PlayerPrompt)

	numberOfPlayers, _ := strconv.Atoi(strings.Trim(c.readLine(), "\n"))

	c.game.Start(numberOfPlayers)

	userInput := c.readLine()

	c.game.Over(extractWinner(userInput))
}

func (c *CLI) readLine() string {
	c.in.Scan()
	return c.in.Text()
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}
