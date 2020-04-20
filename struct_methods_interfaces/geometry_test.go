package struct_methods_interfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.00, 8.00}
	got := Perimeter(rectangle)
	want := 36.00

	if got != want {
		t.Errorf("got %.2f, expected %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{9.00, 6.00}
	got := Area(rectangle)
	want := 54.00

	if got != want {
		t.Errorf("got %.2f, expected %.2f", got, want)
	}
}