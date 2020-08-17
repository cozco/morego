package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk bool
}

func main() {

	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   27,
		},
		ltk: false,
	}

	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
		age:   30,
	}

	p3 := struct {
		first string
		last  string
		age   int
	}{
		first: "Dr",
		last:  "No",
		age:   55,
	}

	fmt.Println(sa1)
	fmt.Println(p2)
	fmt.Println(p3)
}
