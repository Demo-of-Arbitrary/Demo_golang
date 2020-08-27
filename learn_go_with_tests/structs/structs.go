package structs

type Rectangle struct {
	Width  float64
	Height float64
}

func Perimeter(r Rectangle) float64 {
	return (r.Width + r.Height) * 2
}

func Area(r Rectangle) float64 {
	return (r.Width * r.Height)
}
