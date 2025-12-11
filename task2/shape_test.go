package task2

import (
	"fmt"
	"testing"
)

func TestRectangle(t *testing.T) {
	var shape Shape
	shape = Rectangle{10, 2}
	fmt.Printf("The area of the rectangle is: %f\n", shape.Area())
	fmt.Printf("The perimeter of the rectangle is: %f\n", shape.Perimeter())
}

func TestCircle(t *testing.T) {
	var shape Shape
	shape = Circle{5}
	fmt.Printf("The area of the circle is: %f\n", shape.Area())
	fmt.Printf("The perimeter of the circle is: %f\n", shape.Perimeter())
}
