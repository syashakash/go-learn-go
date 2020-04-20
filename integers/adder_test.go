package integers

import "testing"

func TestAdder(t *testing.T) {
	sum := Add(100, 25)
	expected := 125

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}