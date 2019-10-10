package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "hello, "
const spanishHelloPrefix = "hola, "
const frenchHelloPrefix = "bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingsPrefix(language) + name
}

func greetingsPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix

	case french:
		prefix = frenchHelloPrefix

	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
