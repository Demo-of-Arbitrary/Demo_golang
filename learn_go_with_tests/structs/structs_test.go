package structs

import "testing"

func TestPerimeter(t *testing.T) {
	actual := Perimeter(8.0, 2.0)
	expected := 20.0
	if actual != expected {
		t.Errorf("expected %.2f, got %.2f", expected, actual)
	}
}

func TestArea(t *testing.T) {
	actual := Area(8.0, 2.0)
	expected := 16.0
	if actual != expected {
		t.Errorf("expected %.2f, got %.2f", expected, actual)
	}
}
