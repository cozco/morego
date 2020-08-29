package main

import (
	"fmt"
)

func main() {

	ms := "Cosmos"

	strRev(ms)
}

func strRev(s string) string {
	var rs string
	for _, v := range s {
		rs = string(v) + rs
	}
	fmt.Println(rs)
	return rs
}
