package main

import (
	"fmt"
	"sort"
)

type Person struct {
	First string
	Age   int
}

type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
	p1 := Person{"James", 32}
	p2 := Person{"Moneypenny", 27}
	p3 := Person{"Q", 64}
	p4 := Person{"M", 56}

	people := []Person{p1, p2, p3, p4}
	ptest := ByAge{p3, p1, p4, p2}

	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)

	fmt.Println(ptest, "- unsorted")
	ptest.sortName()

}

func (a ByAge) sortName() {
	var ss []string
	for _, v := range a {
		s := fmt.Sprintf("%s%d", v.First+" ", v.Age)
		ss = append(ss, s)
		sort.Strings(ss)
	}

	fmt.Println(ss)
}
