package main

import "fmt"

const hiEngConst = "Hello, "

func main() {
	fmt.Println(hello("", ""))
}

func hello(name, lang string) string {
	if name == "" {
		name = "World"
	}
	if lang == "Spanish" {
		return "Hola, " + name
	}
	return hiEngConst + name
}
