package main

import "fmt"

func main() {

	switch {
	case 2 == 2:
		fmt.Println("test of true")
	case 3 < 1:
		fmt.Println("test of false")
	case 4 == 4:
		fmt.Println("test of fallthrough")
	}
}
