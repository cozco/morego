package main

import "fmt"

func main() {

	for i := 10; i <= 100; i++ {
		x := i % 4
		fmt.Printf("%v\t%v\n", i, x)
	}
}
