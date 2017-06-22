package main

import (
	"fmt"
	"math"
)

func main() {

	rect := Rectangle{2, 100}
	circ := Circle{10}
	box1 := Box{2, 100}

	fmt.Println("Rectangle Area =", getArea(rect))
	fmt.Println("Circle Area =", getArea(circ))
	fmt.Println("Box Area =", getArea(box1))
}

// Shape interface ... feels kinda lika a duct
type Shape interface {
	area() float64
}

// Rectangle struct
type Rectangle struct {
	height float64
	width  float64
}

// Circle struct
type Circle struct {
	radius float64
}

// Box struct
type Box struct {
	height float64
	width  float64
}

func (r Rectangle) area() float64 {
	return r.height * r.width
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (b Box) area() float64 {
	return b.height * b.width * 1000
}

func getArea(shape Shape) float64 {
	fmt.Println(shape)
	return shape.area()
}
