package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	base   float64
	length float64
}

type square struct {
	sideLength float64
}

func main() {
	t := triangle{base: 2, length: 3}
	printArea(t)

	s := square{sideLength: 5}
	printArea(s)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.length
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}
