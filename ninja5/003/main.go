package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	c1 := truck{
		vehicle: vehicle{
			doors: 4,
			color: "red",
		},
		fourWheel: true,
	}

	c2 := sedan{
		vehicle: vehicle{
			doors: 2,
			color: "green",
		},
		luxury: false,
	}

	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c1.color)
	fmt.Println(c2.color)
}
