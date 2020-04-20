package main

import "testing"

func TestSayHelloCustom(t *testing.T)  {
	t.Run("saying hello to people", func(t *testing.T) {
		expected := "Hello, Akash!!"
		got := sayHelloCustom("Akash")

		if got != expected {
			t.Errorf("Got %q, while expected %q", got, expected)
		}
	})

	t.Run("say 'Hello, World!!'", func(t *testing.T) {
		expected := "Hello, World!!"
		got := sayHelloWorld()

		if got != expected {
			t.Errorf("Got %q, while expected %q", got, expected)
		}
	})
}
