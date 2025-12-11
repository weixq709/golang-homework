package task2

import "math"

type Shape interface {
	Area() float32
	Perimeter() float32
}

type Rectangle struct {
	Width  float32
	Height float32
}

type Circle struct {
	Radius float32
}

func (r Rectangle) Area() float32 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float32 {
	return 2 * (r.Width + r.Height)
}

func (c Circle) Area() float32 {
	return c.Radius * c.Radius * math.Pi
}

func (c Circle) Perimeter() float32 {
	return 2 * (c.Radius + c.Radius)
}
