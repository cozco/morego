package main

import "fmt"

func main() {
	n := map[string][]string{
		"bond_james":      {`Shaken, not stirred`, `Martinis`, `Women`},
		"moneypenny_miss": {`moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`},
		"no_dr":           {`Being evil`, `Ice cream`, `Sunsets`},
	}

	n["cosmos"] = []string{"CiCi", "GG", "St. Ives"}

	for i, v := range n {
		for ii, vv := range v {
			fmt.Println(i, ii, vv)
		}
	}
}
