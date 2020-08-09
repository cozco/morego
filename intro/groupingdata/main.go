package main

import "fmt"

func main() {
	// x := make([]int, 10, 100)
	//[]int{4, 5, 6, 7, 8, 9}
	// for i := 0; i < (len(x)); i++ {
	// 	fmt.Println(i, x[i])
	// }

	// y := []int{13, 14, 15, 16}
	// x = append(x, y...)
	// z := append(x, 10, 11, 12)
	// fmt.Println(x)

	// z = append(z[:2], z[4:]...)

	// a := []int{4, 5, 6, 7, 8, 9}
	// b := []int{13, 14, 15, 16}
	// c := [][]int{a, b}

	// fmt.Println(c)

	m := map[string]int{
		"James":      32,
		"Moneypenny": 27,
	}
	fmt.Println(m)
	fmt.Println(m["James"])

	m["Todd"] = 33

	// v, ok := m["Barnabas"]

	// fmt.Println(v, ok)

	// if v, ok := m["James"]; ok {
	// 	fmt.Println(v)
	// }

	for k, v := range m {
		fmt.Println(k, v)
	}

	delete(m, "Todd")
	fmt.Println(m)

}
