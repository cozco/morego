package main

import "fmt"

func main() {

	// n := factorial(4)
	// fmt.Println(n)

	t := facLoop(4)
	fmt.Println(t)
}

func factorial(n int) int {

	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func facLoop(n int) int {
	total := 1

	for ; n > 0; n-- {
		total *= n
	}

	return total

}
