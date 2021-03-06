package poker

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Game interface {
	Start(n int, alertsDestination io.Writer)
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
const BadPlayerInputErrMsg = "Bad value received for number of players, please try again with a number"

func (c *CLI) PlayPoker() {
	fmt.Fprint(c.out, PlayerPrompt)

	numberOfPlayers, err := strconv.Atoi(strings.Trim(c.readLine(), "\n"))

	if err != nil {
		fmt.Fprint(c.out, BadPlayerInputErrMsg)
		return
	}

	c.game.Start(numberOfPlayers, c.out)

	userInput := c.readLine()

	if strings.Contains(userInput, "wins") {
		c.game.Over(extractWinner(userInput))
	}
}

func (c *CLI) readLine() string {
	c.in.Scan()
	return c.in.Text()
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}
