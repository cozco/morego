package main

import "fmt"

type person struct {
	fName string
	lNmae string
}

type secretAgent struct {
	person
	ltk bool
}

type human interface {
	speak()
}

func humanPrinter(h human) {
	fmt.Println("I love humans like", h)
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.lNmae, s.fName, s.lNmae)
}

func main() {

	p1 := secretAgent{
		person: person{
			fName: "James",
			lNmae: "Bond",
		},
		ltk: true,
	}

	n1 := 50 + 11

	func(n int) {
		fmt.Println("i'm adding these two numbers", n)
	}(n1)

	n2 := 27

	g := func(n int) {
		fmt.Println("i'm adding these two numbers", n)
	}

	g(n2)

	humanPrinter(p1)

	p1.speak()

}
