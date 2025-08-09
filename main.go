package main

import "fmt"

const (
	spanish      = "Spanish"
	french       = "French"
	engHello     = "Hello, "
	spanishHello = "Hola, "
	frenchHello  = "Bonjour, "
)

func main() {
	fmt.Println(hello("", ""))
}

func hello(name, lang string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(lang) + name
}

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case spanish:
		prefix = spanishHello
	case french:
		prefix = frenchHello
	default:
		prefix = engHello
	}

	return prefix
}
