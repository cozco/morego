package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {

	p1 := person{
		first: "Cozco",
		last:  "Onana",
		age:   27,
	}

	p1.speak()

}

func (p person) speak() {
	fmt.Printf("My name is %v and my age is %v", (p.first + " " + p.last), p.age)
}
