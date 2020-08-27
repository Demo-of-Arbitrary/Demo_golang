package structs

import "testing"

func TestPerimeter(t *testing.T) {
	reactAngle := Rectangle{8.0, 2.0}
	actual := Perimeter(reactAngle)
	expected := 20.0
	if actual != expected {
		t.Errorf("expected %.2f, got %.2f", expected, actual)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, name string, shape Shape, expected float64) {
		t.Helper()
		actual := shape.Area()
		t.Run(name, func(t *testing.T) {
			if actual != expected {
				t.Errorf("%#v expected %g, got %g", shape, expected, actual)
			}
		})
	}

	tests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "RectAngle", shape: Rectangle{Width: 8.0, Height: 2.0}, hasArea: 16.0},
		{name: "Circle", shape: Circle{Radius: 10.0}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{A: 3.0, B: 4.0, C: 5.0}, hasArea: 6.0},
	}

	for _, tt := range tests {
		checkArea(t, tt.name, tt.shape, tt.hasArea)
	}
}
