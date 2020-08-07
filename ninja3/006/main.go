package main

import "fmt"

func main() {
	x := 10

	if x == 9 {
		fmt.Println("x is equal to 9")
	} else if x > 15 {
		fmt.Printf("%v is greater than 5", x)
	} else {
		fmt.Println("No luck", x)
	}
}
