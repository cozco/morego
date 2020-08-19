package main

import (
	"fmt"
)

type add int

func main() {
	var a add = 3
	// var b add = 3
	// var c add = 4

	b := a.adder()
	c := a.adder()

	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(c())
	fmt.Println(c())

}

func (a add) adder() func() int {
	var x int

	return func() int {
		x++
		return x * int(a)
	}
}
