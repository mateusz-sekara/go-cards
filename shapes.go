package main

import "fmt"

type shape interface {
	getArea() float64
}

func printArea(s shape) {
	fmt.Println("Shape's area:", s.getArea())
}

type triangle struct {
	height float64
	base   float64
}

func (t triangle) getArea() float64 {
	return t.height * t.base * 0.5
}

type square struct {
	sideLength float64
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func shapesSandbox() {
	t := triangle{3.2, 6.1}
	s := square{2.5}

	printArea(t)
	printArea(s)
}
