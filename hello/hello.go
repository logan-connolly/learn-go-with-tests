package main

import "fmt"

const spanish = "Spanish"
const french = "French"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) string {
	switch language {
	case spanish:
		return "Hola, "
	case french:
		return "Bonjour, "
	default:
		return "Hello, "
	}
}

func main() {
	fmt.Println(Hello("", ""))
}
