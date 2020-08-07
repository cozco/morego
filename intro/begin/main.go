package main

import "fmt"

type cosmos int

func main() {
	// y := 42
	// z := `"My name is wicki slicky Slim shady"`
	// fmt.Println(y)
	// fmt.Printf("%T\n", y)
	// fmt.Println(z)
	// fmt.Printf("%T\n", z)

	var b cosmos
	b = 15
	fmt.Println(b)

	a := int(b)
	fmt.Println(a)

}
