package main

import "fmt"

// Shape defines shared behavior.
type Shape interface {
	Area() float64
	Name() string
}

type Rectangle struct {
	W, H float64
}

func (r Rectangle) Area() float64 {
	return r.W * r.H
}

func (r Rectangle) Name() string {
	return "rectangle"
}

type Circle struct {
	R float64
}

func (c Circle) Area() float64 {
	return 3.1415926 * c.R * c.R
}

func (c Circle) Name() string {
	return "circle"
}

func main() {
	// Polymorphism: both types satisfy Shape implicitly.
	shapes := []Shape{
		Rectangle{W: 3, H: 4},
		Circle{R: 2},
	}

	for _, s := range shapes {
		fmt.Println(s.Name(), "area:", s.Area())
	}

	// Type assertion retrieves concrete value when needed.
	var anyShape Shape = Circle{R: 5}
	if c, ok := anyShape.(Circle); ok {
		fmt.Println("type assertion circle radius:", c.R)
	}
}
