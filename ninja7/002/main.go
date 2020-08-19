package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

func changeMe(p *person) {
	(*p).first = "Cozco"
	(*p).last = "The Great"
	fmt.Println(p.first, p.last)
}

func main() {

	p1 := person{
		first: "Cici",
		last:  "Smiles",
	}

	fmt.Println(p1.first, p1.last)

	changeMe(&p1)

}
