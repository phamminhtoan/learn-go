package geometry2

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width float64
	Height float64
}

type Circles struct {
	Radius float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2*(rectangle.Width + rectangle.Height)
}

func Area(rectangle Rectangle) float64  {
	return  rectangle.Width * rectangle.Height
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circles) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}