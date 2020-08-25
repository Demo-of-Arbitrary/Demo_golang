package main

import "fmt"

const helloEnglishPrefix = "Hello "
const helloSpanishPrefix = "Hola "

func Hello(name string, lang string) string {
	helloPrefix := helloEnglishPrefix
	if lang == "spanish" {
		helloPrefix = helloSpanishPrefix
	}
	if name == "" {
		name = "World"
	}
	return helloPrefix + name
}
func main() {
	fmt.Println(Hello("World", ""))
}
