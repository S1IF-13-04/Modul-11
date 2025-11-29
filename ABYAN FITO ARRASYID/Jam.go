package main

import "fmt"

func main() {
	var jam int
	fmt.Scan(&jam)

	switch {
	case jam == 0:
		fmt.Println("12 AM")
	case jam < 12:
		fmt.Println(jam, "AM")
	case jam == 12:
		fmt.Println("12 PM")
	default:
		fmt.Println(jam-12, "PM")
	}
}
