package shapes

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

func (rect Rectangle) Perimeter() float64 {
	return 2 * (rect.Width + rect.Height)
}

func (rect Rectangle) Area() float64 {
	return rect.Width * rect.Height
}

func (cir Circle) Area() float64 {
	return math.Pi * cir.Radius * cir.Radius
}

func (cir Circle) Perimeter() float64 {
	return 2 * math.Pi * cir.Radius
}

func (tri Triangle) Perimeter() float64 {
	return tri.Base + 2*tri.Height
}

func (tri Triangle) Area() float64 {
	return tri.Base / 2 * tri.Height
}
