package main

import "fmt"

const englishHelloPrefix = "Hello, "

// Hello ...
// Say hello to someone or the world
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Jacob"))
}
