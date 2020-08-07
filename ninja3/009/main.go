package main

import "fmt"

func main() {
	favSport := "Handegg"

	switch favSport {
	case "Hockey":
		fmt.Println("I saw a game in Detroit once")
	case "Football", "Handegg":
		fmt.Println("Mahomes!!!")
	case "Basketball":
		fmt.Println("Team Lebron")
	case "Baseball":
		fmt.Println("Err snooze")
	case "Soccer":
		fmt.Println("Arsenal is my club de futbol")
	default:
		fmt.Println("You don't watch sports? Yo we can't be friends")
	}
}
