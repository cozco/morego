package main

import "fmt"

type person struct {
	first    string
	last     string
	iceCream []string
}

func main() {

	// p1 := person{
	// 	first:    "Cozco",
	// 	last:     "Oman",
	// 	iceCream: []string{"chocolate", "strawberry", "lemon"},
	// }

	// p2 := person{
	// 	first:    "Cici",
	// 	last:     "Coman",
	// 	iceCream: []string{"chocolate", "strawberry", "vanilla"},
	// }

	p3 := map[string]person{
		"Oman": person{
			first:    "Cozco",
			last:     "Oman",
			iceCream: []string{"chocolate", "strawberry", "lemon"},
		},
		"Coman": person{
			first:    "Cici",
			last:     "Coman",
			iceCream: []string{"chocolate", "strawberry", "vanilla"},
		},
	}

	// fmt.Println(p1.first, p1.last)
	// for _, v := range p1.iceCream {
	// 	fmt.Println(v)
	// }

	// fmt.Println(p2.first, p2.last)
	// for _, v := range p2.iceCream {
	// 	fmt.Println(v)
	// }

	for _, v := range p3 {
		fmt.Println(v.first)
		fmt.Println(v.last)
		for _, x := range v.iceCream {
			fmt.Println(x)
		}
	}

}
