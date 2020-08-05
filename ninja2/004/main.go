package main

import "fmt"

func main() {
	x := 10
	fmt.Printf("%d\t%b\t%#U", x, x, x)
	y := x << 10
	fmt.Println()
	fmt.Printf("%d\t%b\t%#X", y, y, y)
}
