package main

import "fmt"

type triangle struct {
	high float64
	base float64
}

type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func main() {
	triangle := triangle{12.2, 12.2}
	square := square{12.2}

	printArea(triangle)
	printArea(square)

}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return t.base * t.high * 0.5
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
