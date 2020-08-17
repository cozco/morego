package main

import (
	"fmt"
)

func main() {
	s, l := 6, 6
	a := areaCaller()
	b := a(s, l)
	fmt.Println(b)
}

func areaCaller() func(l int, w int) int {
	return func(l int, w int) int {
		a := l * w
		return a
	}
}
