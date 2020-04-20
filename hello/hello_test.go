package main

import "testing"

func TestSayHelloCustom(t *testing.T)  {

	assertCorrectMessage := func(t *testing.T, got, expected string) {
		t.Helper()
		if got != expected {
			t.Errorf("Got %q, while expected %q", got, expected)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		expected := "Hello, Akash!!"
		got := sayHello("Akash", "EN")
		assertCorrectMessage(t, got, expected)
	})

	t.Run("say 'Hello, World!!'", func(t *testing.T) {
		expected := "Hello, World!!"
		got := sayHello("", "EN")
		assertCorrectMessage(t, got, expected)
	})

	t.Run("in Spanish", func(t *testing.T) {
		expected := "Hola, Elodie!!"
		got := sayHello("Elodie", "ES")
		assertCorrectMessage(t, got, expected)
	})

	t.Run("in French",func(t *testing.T) {
		expected := "Bonjour, Elodie!!"
		got := sayHello("Elodie", "FR")
		assertCorrectMessage(t, got, expected)
	})
}
