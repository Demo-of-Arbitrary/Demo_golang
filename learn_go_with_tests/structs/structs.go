package structs

import (
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return (r.Width * r.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return (math.Pi * c.Radius * c.Radius)
}

type Triangle struct {
	A float64
	B float64
	C float64
}

func (t Triangle) Area() float64 {
	a := t.A
	b := t.B
	c := t.C
	p := (a + b + c) / 2
	return math.Sqrt(p*(p-a) + p*(p-b) + p*(p-c))
}

func Perimeter(r Rectangle) float64 {
	return (r.Width + r.Height) * 2
}
