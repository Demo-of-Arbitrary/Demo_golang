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
	reactAngle := Rectangle{8.0, 2.0}
	actual := Area(reactAngle)
	expected := 16.0
	if actual != expected {
		t.Errorf("expected %.2f, got %.2f", expected, actual)
	}
}
