package main

import "fmt"

type shape interface {
	getArea() float64
	getType() string
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

func main() {
	myTriangle := triangle{height: 3, base: 5}
	mySquare := square{sideLength: 5}

	printArea(myTriangle)
	printArea(mySquare)
}

func (t triangle) getArea() float64 {
	return (t.base * t.height) / 2
}

func (t triangle) getType() string {
	return "triangle"
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (s square) getType() string {
	return "square"
}

func printArea(s shape) {
	fmt.Println("Area for:", s.getType(), s.getArea())
}
