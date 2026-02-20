package sctuctsmethodsinterfaces

import "math"

type Rectangle struct {
	Width  float64
	Length float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Length
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func Perimiter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Length)
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return (.5 * t.Base) * t.Height
}
