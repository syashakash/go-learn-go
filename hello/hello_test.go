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
		got := SayHello("Akash", "EN")
		assertCorrectMessage(t, got, expected)
	})

	t.Run("say 'Hello, World!!'", func(t *testing.T) {
		expected := "Hello, World!!"
		got := SayHello("", "EN")
		assertCorrectMessage(t, got, expected)
	})

	t.Run("in Spanish", func(t *testing.T) {
		expected := "Hola, Elodie!!"
		got := SayHello("Elodie", "ES")
		assertCorrectMessage(t, got, expected)
	})

	t.Run("in French",func(t *testing.T) {
		expected := "Bonjour, Elodie!!"
		got := SayHello("Elodie", "FR")
		assertCorrectMessage(t, got, expected)
	})
}
