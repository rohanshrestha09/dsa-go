package main

import "fmt"

type Rectangle struct {
	length  float32
	breadth float32
}

func (r Rectangle) Area() {
	fmt.Print("Area of rectangle is ", r.length*r.breadth)
}

func Run() {
	rectangle := Rectangle{length: 10, breadth: 20}

	rectangle.Area()
}
