package main

import "fmt"

type shape interface{
	area() float64
}
type triangle struct{
	height float64
	base float64
}
type square struct{
	sideLength float64
}

func main() {
	s := square{
		sideLength: 6,
	}
	t := triangle{
		height: 6,
		base: 2,
	}

	printArea(s) 
	printArea(t) 
}

func (t triangle) area() float64 {
	return t.height * t.base * 0.5
}

func (s square) area() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println(s.area())
}
