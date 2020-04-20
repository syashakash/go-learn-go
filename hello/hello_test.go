package main

import "testing"

func  TestSayHello(t *testing.T)  {
	expected := "Hello, World!!"
	got := sayHello()

	if got != expected {
		t.Errorf("Got %q, while expected %q", got, expected)
	}
}