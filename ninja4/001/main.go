package main

import "fmt"

func main() {
	var x = [5]int{1, 2, 3, 4, 5}

	for _, v := range x {
		fmt.Println(v)
	}

	fmt.Printf("%T\t%v", x, x)
}
