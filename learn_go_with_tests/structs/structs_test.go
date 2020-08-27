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
	t.Run("rectangles", func(t *testing.T) {
		rectAngle := Rectangle{8.0, 2.0}
		actual := rectAngle.Area()
		expected := 16.0
		if actual != expected {
			t.Errorf("expected %.2f, got %.2f", expected, actual)
		}
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		actual := circle.Area()
		expected := 314.1592653589793

		if actual != expected {
			t.Errorf("expected %g, got %g", expected, actual)
		}
	})
}
