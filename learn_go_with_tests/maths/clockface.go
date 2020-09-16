package clockface

import "time"

type Point struct {
	X, Y float64
}
type clockface struct {
	Point Point
}

func SecondHand(t time.Time) Point {
	return Point{150, 60}
}
