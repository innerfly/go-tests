package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const engHello = "Hello, "
const spanishHello = "Hola, "
const frenchHello = "Bonjour, "

func main() {
	fmt.Println(hello("", ""))
}

func hello(name, lang string) string {
	if name == "" {
		name = "World"
	}

	prefix := engHello
	switch lang {
	case spanish:
		prefix = spanishHello
	case french:
		prefix = frenchHello
	}

	return prefix + name
}
