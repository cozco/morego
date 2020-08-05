package main

import "fmt"

func main() {
	a := 12 == 10
	b := 2 <= 5
	c := 67 >= 12
	j := 45 != 30
	g := 12 < 345
	i := 23.21 > 4.56

	fmt.Println(a, b, c, j, g, i)
}
