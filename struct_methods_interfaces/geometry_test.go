package struct_methods_interfaces

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.00, 8.00)
	want := 36.00

	if got != want {
		t.Errorf("got %.2f, expected %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	got := Area(9.0, 6.0)
	want := 54.00
	if got != want {
		t.Errorf("got %.2f, expected %.2f", got, want)
	}
}