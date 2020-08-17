package main

import "fmt"

type person struct {
	first    string
	last     string
	iceCream []string
}

func main() {

	p1 := person{
		first:    "Cozco",
		last:     "Oman",
		iceCream: []string{"chocolate", "strawberry", "lemon"},
	}

	p2 := person{
		first:    "Cici",
		last:     "Oman",
		iceCream: []string{"chocolate", "strawberry", "vanilla"},
	}

	fmt.Println(p1.first, p1.last)
	for _, v := range p1.iceCream {
		fmt.Println(v)
	}

	fmt.Println(p2.first, p2.last)
	for _, v := range p2.iceCream {
		fmt.Println(v)
	}

}
