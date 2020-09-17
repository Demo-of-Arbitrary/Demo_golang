package main

import (
	"os"
	"time"

	clockface "github.com/Demo_golang/learn_go_with_tests/maths"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
