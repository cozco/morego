package main

import "fmt"

func main() {
	n := foo(5)
	fmt.Println(n)
	s := bar("Yo Wassup")
	fmt.Println(s)

}

func foo(n int) int {
	return n

}

func bar(s string) string {
	return s

}
