package main

import "fmt"

func main() {

	a := 10

	b := func(x int) int {
		b := x * 2
		fmt.Println(b)
		return b
	}

	b(a)

}
