package main

import "fmt"

type cosmos int

var x cosmos

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
