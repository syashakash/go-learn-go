package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanish = "ES"
const english = "EN"
const french = "FR"
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func SayHello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name + "!!"
}

func greetingPrefix(language string) (prefix string) {
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
func main() {
	fmt.Println(sayHello("", "ES"))
}
