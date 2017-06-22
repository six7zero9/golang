package main

import "fmt"

func main() {

	rect1 := Rectangle{leftX: 0, topY: 50, height: 10, width: 10}

	fmt.Println("Rectangle is", rect1.width, "wide.")
	fmt.Println("Rectangle's Area =", rect1.area())

}

// Rectangle struct generic comment
type Rectangle struct {
	leftX  float64
	topY   float64
	height float64
	width  float64
}

func (rect Rectangle) area() float64 {
	fmt.Println(rect)
	return rect.width * rect.height
}
