package main

import "fmt"

func main() {
	fmt.Println(hello(""))
}

const hiConst = "Hello, "

func hello(name string) string {
	if name == "" {
		name = "World"
	}
	return hiConst + name
}
