package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First string
	Last  string
	Age   int
}

type uperson struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	js := `[{"First":"Cozco","Last":"Oman","Age":27},{"First":"Cici","Last":"Oman","Age":31}]`

	bs := []byte(js)

	var p []uperson

	err := json.Unmarshal(bs, &p)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(p)
	fmt.Printf("%+v", p)
	fmt.Printf("\n%T\t%T", js, bs)
}

func jsMarshal() {

	p1 := person{
		First: "Cozco",
		Last:  "Oman",
		Age:   27,
	}

	p2 := person{
		First: "Cici",
		Last:  "Oman",
		Age:   31,
	}

	p := []person{p1, p2}

	bs, err := json.Marshal(p)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(p)
	fmt.Println(string(bs))

}
