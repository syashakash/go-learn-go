package main

import "fmt"

const englishHelloPrefix = "Hello, "

func sayHelloWorld() string {
	return englishHelloPrefix + "World!!"
}

func sayHelloCustom(name string) string {
	return englishHelloPrefix + name + "!!"
}

func main() {
	fmt.Println(sayHelloWorld())
	fmt.Println(sayHelloCustom("Akash"))
}
