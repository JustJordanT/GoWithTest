package main

import "fmt"

const spanish = "Spanish"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const french = "French"

func main() {
	fmt.Println(Hello("JT", "English"))
}

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return languagePrefix(language) + name

}

func languagePrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
