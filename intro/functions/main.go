package main

import "fmt"

func main() {

	xi := []int{2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(sum(xi...))

	//foo(2, 3, 4, 5, 6, 7, 8, 9)
}

func foo(x ...int) {
	fmt.Println(x)
	fmt.Printf("%T\t", x)

}

func sum(x ...int) int {

	sum := 0
	for _, v := range x {
		sum += v
		//fmt.Println(v)
	}
	return sum
}
