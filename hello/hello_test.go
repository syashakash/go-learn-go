package main

import "testing"

func TestSayHelloCustom(t *testing.T)  {
	expected := "Hello, Akash!!"
	got := sayHelloCustom("Akash")

	if got != expected {
		t.Errorf("Got %q, while expected %q", got, expected)
	}
}

func TestSayHelloWorld(t *testing.T) {

	expected := "Hello, World!!"
	got := sayHelloWorld()

	if got != expected {
		t.Errorf("Got %q, while expected %q", got, expected)
	}
}