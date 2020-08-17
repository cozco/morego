package main

import (
	"fmt"
)

func main() {
	xi := []int{7, 7}

	areaCaller(area, xi...)

}

func areaCaller(f func(l ...int) int, xi ...int) int {
	return f(xi...)
}

func area(x ...int) int {
	a := x[0] * x[1]
	fmt.Println(a)
	return a
}
