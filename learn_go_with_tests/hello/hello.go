package main

import "fmt"

func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}
	return prefix(lang) + name
}
func prefix(lang string) (helloPrefix string) {
	const helloEnglishPrefix = "Hello "
	const helloSpanishPrefix = "Hola "
	const helloFrenchPrefix = "Bonjour "

	switch lang {
	case "spanish":
		helloPrefix = helloSpanishPrefix
	case "french":
		helloPrefix = helloFrenchPrefix
	default:
		helloPrefix = helloEnglishPrefix
	}
	return
}
func main() {
	fmt.Println(Hello("World", ""))
}
