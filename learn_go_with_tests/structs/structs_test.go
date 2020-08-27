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
	checkArea := func(t *testing.T, shape Shape, expected float64) {
		t.Helper()
		actual := shape.Area()
		if actual != expected {
			t.Errorf("expected %f, got %f", expected, actual)
		}
	}
	t.Run("rectangles", func(t *testing.T) {
		rectAngle := Rectangle{8.0, 2.0}
		expected := 16.0
		checkArea(t, rectAngle, expected)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		expected := 314.1592653589793

		checkArea(t, circle, expected)
	})
}
