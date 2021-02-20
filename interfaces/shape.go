package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}

type Rectange struct {
	width, height float64
}

func getArea(s Shape) float64 {
	return s.area()
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectange) area() float64 {
	return r.width * r.height
}

func main() {
	circle := Circle { 0, 0, 5}
	rectangle := Rectange { 10, 5}

	fmt.Printf("Circle Area: %f\n", getArea(circle))
	fmt.Printf("Rectangle Area: %f\n", getArea(rectangle))
}