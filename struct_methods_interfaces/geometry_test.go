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
	areaTests := []struct {
		name string
		shape Shape
		hasArea  float64
	}{
		{name:"Rectangle", shape:Rectangle{12, 6}, hasArea:72.0},
		{name:"Circle", shape:Circle{10}, hasArea:314.1592653589793},
		{name:"Triangle", shape:Triangle{12, 8}, hasArea:36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.hasArea {
			t.Errorf("%#v got %g want %g", tt.name, got, tt.hasArea)
		}
	}
}