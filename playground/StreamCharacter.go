package main

import (
	"fmt"
)

type StreamChecker struct {
	words []string
}

func Constructor(words []string) StreamChecker {
	sc := StreamChecker{words}
	return sc
}

func (this *StreamChecker) Query(letter byte) bool {
	return false
}
func main() {
	var sArr [2]string
	sArr[0] = "abc"
	sArr[1] = "def"
	m := Constructor(sArr)
	fmt.Println(m)
}
