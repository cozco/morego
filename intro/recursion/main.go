package main

import "fmt"

func main() {

	n := factorial(4)
	fmt.Println(n)

	// t := facLoop(4)
	// fmt.Println(t)
}

func factorial(n int) int {

	if n == 0 {
		return 1
	}

	a := n * factorial(n-1)
	fmt.Println(a)
	return a
}

func facLoop(n int) int {
	total := 1

	for ; n > 0; n-- {
		fmt.Println(n, total)
		total *= n
	}

	return total

}
