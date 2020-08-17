package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4}
	y := foo(x...)
	fmt.Println(y)

	t := bar(x)
	fmt.Println(t)
}

func foo(x ...int) int {

	var y int

	for _, v := range x {
		y += v
	}

	return y
}

func bar(x []int) int {

	y := foo(x...)

	return y

}
