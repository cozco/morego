package main

import "fmt"

func main() {

	a := 10

	b := func(x int) int {
		c := x * 2
		fmt.Println(c)
		return c
	}

	b(a)
	fmt.Printf("%T", b)

}
