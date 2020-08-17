package main

import (
	"fmt"
)

func main() {

	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	t := evenSum(sum, xi...)
	fmt.Println(t)

	o := oddSum(sum, xi...)
	fmt.Println(o)
}

func sum(x ...int) int {
	n := 0
	for _, v := range x {
		n += v
	}
	return n
}

func evenSum(f func(x ...int) int, y ...int) int {
	var xi []int
	for _, v := range y {
		if v%2 == 0 {
			xi = append(xi, v)
		}
	}
	total := f(xi...)
	return total
}

func oddSum(f func(x ...int) int, y ...int) int {
	xi := []int{}

	for _, v := range y {
		if v%2 != 0 {
			xi = append(xi, v)
		}
	}
	return f(xi...)
}
