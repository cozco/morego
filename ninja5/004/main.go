package main

import "fmt"

func main() {

	t1 := struct {
		name  string
		sport string
		year  int
	}{
		name:  "Arsenal",
		sport: "Soccer",
		year:  1886,
	}

	fmt.Println(t1)
}
