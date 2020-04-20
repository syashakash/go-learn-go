package main

import "fmt"

func sayHelloWorld() string {
	return "Hello, World!!"
}

func sayHelloCustom(name string) string {
	return "Hello, " + name + "!!"
}

func main() {
	fmt.Println(sayHelloWorld())
	fmt.Println(sayHelloCustom("Akash"))
}
