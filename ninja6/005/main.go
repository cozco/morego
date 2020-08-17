package main

import "fmt"

type square struct {
	length int
	width  int
}

type circle struct {
	radius float64
}

type shape interface {
	area() int
}

func main() {
	s1 := square{
		length: 6,
		width:  6,
	}

	c1 := circle{
		radius: 12.2,
	}

	a := info(s1)
	b := info(c1)

	fmt.Println(a, b)
}

func (s square) area() int {
	a := s.length * s.width
	return a
}

func (c circle) area() int {
	a := 3.14 * (c.radius * c.radius)
	return int(a)
}

func info(s shape) int {

	var x int

	switch s.(type) {
	case circle:
		x = s.(circle).area()
	case square:
		x = s.(square).area()
	}

	//return s.area (The Go compiler knows that you're referencing a circle or square)

	return x

}
