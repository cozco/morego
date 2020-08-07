// package main

// import "fmt"

// func main() {

// 	for i := 65; i <= 90; i++ {
// 		fmt.Println(i)
// 		fmt.Printf("\t%#U\n\t%#U\n\t%#U\n", i, i, i)
// 	}
// }

package main

import "fmt"

func main() {

	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for x := 0; x < 3; x++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}
